package runtime

import (
	"context"
	"github.com/autom8ter/graphik/api/model"
	"net/http"
)

const (
	authCtxKey = "x-graphik-auth-ctx"
)

func (a *Runtime) toContext(r *http.Request, payload map[string]interface{}) *http.Request {
	path := model.Path{
		Type: "user",
		ID:   payload["sub"].(string),
	}
	n, ok := a.nodes.Get(path)
	if !ok {
		n = a.nodes.Set(&model.Node{
			Path:       path,
			Attributes: n.Attributes,
		})
	}
	return r.WithContext(context.WithValue(r.Context(), authCtxKey, n))
}

func (s *Runtime) GetNode(ctx context.Context) *model.Node {
	val, ok := ctx.Value(authCtxKey).(*model.Node)
	if !ok {
		return nil
	}
	return val
}
