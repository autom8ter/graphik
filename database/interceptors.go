package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/graphik/gen/go"
	"github.com/autom8ter/graphik/helpers"
	"github.com/autom8ter/graphik/logger"
	"github.com/google/cel-go/cel"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	authCtxKey   = "x-graphik-auth-ctx"
	identityType = "identity"
	emailClaim   = "email"
	methodCtxKey = "x-grpc-full-method"
)

func (g *Graph) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		idToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return nil, err
		}
		idTokenHash := helpers.Hash([]byte(idToken))
		if val, ok := g.jwtCache.Get(idTokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := g.check(ctx, info.FullMethod, req, payload)
			if err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
		payload, err := g.verifyJWT(idToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}
		var exp int64
		if val, ok := payload["exp"].(int64); ok {
			if val < time.Now().Unix() {
				return nil, status.Errorf(codes.Unauthenticated, "token expired")
			}
			exp = val
		}
		if val, ok := payload["exp"].(int); ok {
			if int64(val) < time.Now().Unix() {
				return nil, status.Errorf(codes.Unauthenticated, "token expired")
			}
			exp = int64(val)
		}
		ctx = g.methodToContext(ctx, info.FullMethod)
		g.jwtCache.Set(idTokenHash, payload, time.Unix(exp, 0).Sub(time.Now()))
		ctx, err = g.check(ctx, info.FullMethod, req, payload)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (g *Graph) StreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		idToken, err := grpc_auth.AuthFromMD(ss.Context(), "Bearer")
		if err != nil {
			return err
		}
		idTokenHash := helpers.Hash([]byte(idToken))
		if val, ok := g.jwtCache.Get(idTokenHash); ok {
			payload := val.(map[string]interface{})
			ctx, err := g.check(ss.Context(), info.FullMethod, srv, payload)
			if err != nil {
				return err
			}
			wrapped := grpc_middleware.WrapServerStream(ss)
			wrapped.WrappedContext = ctx
			return handler(srv, wrapped)
		}
		payload, err := g.verifyJWT(idToken)
		if err != nil {
			return status.Errorf(codes.Unauthenticated, err.Error())
		}
		var exp int64
		if val, ok := payload["exp"].(int64); ok {
			if val < time.Now().Unix() {
				return status.Errorf(codes.Unauthenticated, "token expired")
			}
			exp = val
		}
		if val, ok := payload["exp"].(int); ok {
			if int64(val) < time.Now().Unix() {
				return status.Errorf(codes.Unauthenticated, "token expired")
			}
			exp = int64(val)
		}
		g.jwtCache.Set(idTokenHash, payload, time.Unix(exp, 0).Sub(time.Now()))
		ctx, err := g.check(ss.Context(), info.FullMethod, srv, payload)
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}

func (a *Graph) identityToContext(ctx context.Context, payload map[string]interface{}) (context.Context, *apipb.Doc, error) {
	if _, ok := payload["email"].(string); !ok {
		return nil, nil, errors.New("email not present in token claims")
	}
	docs, _ := a.SearchDocs(ctx, &apipb.Filter{
		Gtype:      identityType,
		Expression: fmt.Sprintf(`doc.attributes.email.contains("%s")`, payload["email"].(string)),
		Limit:      1,
	})
	if docs == nil || len(docs.GetDocs()) == 0 {
		docs = &apipb.Docs{}
		logger.Info("creating identity", zap.String("email", payload["email"].(string)))
		strct, err := structpb.NewStruct(payload)
		if err != nil {
			return nil, nil, err
		}
		doc, err := a.createIdentity(ctx, &apipb.DocConstructor{
			Gtype:      identityType,
			Attributes: strct,
		})
		if err != nil {
			return nil, nil, err
		}
		docs.Docs = []*apipb.Doc{doc}
	}
	return context.WithValue(ctx, authCtxKey, docs.GetDocs()[0]), docs.GetDocs()[0], nil
}

func (s *Graph) getIdentity(ctx context.Context) *apipb.Doc {
	val, ok := ctx.Value(authCtxKey).(*apipb.Doc)
	if ok {
		return val
	}
	val2, ok := ctx.Value(authCtxKey).(apipb.Doc)
	if ok {
		return &val2
	}
	return nil
}

func (r *Graph) getMethod(ctx context.Context) string {
	val, ok := ctx.Value(methodCtxKey).(string)
	if ok {
		return val
	}
	return ""
}

func (r *Graph) methodToContext(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, methodCtxKey, path)
}

func (g *Graph) verifyJWT(token string) (map[string]interface{}, error) {
	message, err := jws.ParseString(token)
	if err != nil {
		return nil, err
	}
	g.jwksMu.RLock()
	defer g.jwksMu.RUnlock()
	if g.jwksSet == nil {
		data := map[string]interface{}{}
		if err := json.Unmarshal(message.Payload(), &data); err != nil {
			return nil, err
		}
		return data, nil
	}
	if len(message.Signatures()) == 0 {
		return nil, fmt.Errorf("zero jws signatures")
	}
	kid, ok := message.Signatures()[0].ProtectedHeaders().Get("kid")
	if !ok {
		return nil, fmt.Errorf("jws kid not found")
	}
	algI, ok := message.Signatures()[0].ProtectedHeaders().Get("alg")
	if !ok {
		return nil, fmt.Errorf("jw alg not found")
	}
	alg, ok := algI.(jwa.SignatureAlgorithm)
	if !ok {
		return nil, fmt.Errorf("alg type cast error")
	}
	keys := g.jwksSet.LookupKeyID(kid.(string))
	if len(keys) == 0 {
		return nil, errors.Errorf("failed to lookup kid: %s - zero keys", kid.(string))
	}
	var key interface{}
	if err := keys[0].Raw(&key); err != nil {
		return nil, err
	}
	payload, err := jws.Verify([]byte(token), alg, key)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (g *Graph) check(ctx context.Context, method string, req interface{}, payload map[string]interface{}) (context.Context, error) {
	ctx = g.methodToContext(ctx, method)
	ctx, identity, err := g.identityToContext(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if g.authorizers.Len() > 0 {
		now := time.Now()
		request := &apipb.Request{
			Method:    method,
			Identity:  identity,
			Timestamp: timestamppb.New(now),
		}
		if val, ok := req.(proto.Message); ok {
			bits, _ := helpers.MarshalJSON(val)
			reqMap := map[string]interface{}{}
			if err := json.Unmarshal(bits, &reqMap); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			request.Request = apipb.NewStruct(reqMap)
		}
		var programs []cel.Program
		g.rangeAuthorizers(func(a *authorizer) bool {
			programs = append(programs, a.program)
			return true
		})
		result, err := g.vm.Auth().Eval(request, programs...)
		if err != nil {
			return nil, err
		}
		if !result {
			return nil, status.Error(codes.PermissionDenied, "request authorization = denied")
		}
	}
	if g.getIdentity(ctx) == nil {
		return nil, status.Error(codes.Internal, "empty identity")
	}
	return ctx, nil
}
