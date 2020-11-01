package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/autom8ter/graphik/command"
	"github.com/autom8ter/graphik/generic"
	"github.com/autom8ter/graphik/graph/generated"
	"github.com/autom8ter/graphik/graph/model"
	"github.com/autom8ter/machine"
)

func (r *mutationResolver) CreateNode(ctx context.Context, input model.NodeConstructor) (*model.Node, error) {
	if input.Path.ID == "" {
		random := generic.UUID()
		input.Path.ID = random
	}
	if input.Path.Type == "" {
		input.Path.Type = generic.Default
	}
	res, err := r.store.Execute(&command.Command{
		Op:  command.CREATE_NODE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Node), nil
}

func (r *mutationResolver) PatchNode(ctx context.Context, input *model.Patch) (*model.Node, error) {
	res, err := r.store.Execute(&command.Command{
		Op:  command.PATCH_NODE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Node), nil
}

func (r *mutationResolver) DelNode(ctx context.Context, input model.Path) (*model.Counter, error) {
	res, err := r.store.Execute(&command.Command{
		Op:  command.DELETE_NODE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Counter), nil
}

func (r *mutationResolver) CreateEdge(ctx context.Context, input model.EdgeConstructor) (*model.Edge, error) {
	if input.Path.ID == "" {
		random := generic.UUID()
		input.Path.ID = random
	}
	if input.Path.Type == "" {
		input.Path.Type = generic.Default
	}
	res, err := r.store.Execute(&command.Command{
		Op:  command.CREATE_EDGE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Edge), nil
}

func (r *mutationResolver) PatchEdge(ctx context.Context, input model.Patch) (*model.Edge, error) {
	res, err := r.store.Execute(&command.Command{
		Op:  command.PATCH_EDGE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Edge), nil
}

func (r *mutationResolver) DelEdge(ctx context.Context, input model.Path) (*model.Counter, error) {
	res, err := r.store.Execute(&command.Command{
		Op:  command.DELETE_EDGE,
		Val: input,
	})
	if err != nil {
		return nil, err
	}
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*model.Counter), nil
}

func (r *queryResolver) GetNode(ctx context.Context, input model.Path) (*model.Node, error) {
	return r.store.Node(ctx, input)
}

func (r *queryResolver) GetNodes(ctx context.Context, input model.Filter) ([]*model.Node, error) {
	return r.store.Nodes(ctx, input)
}

func (r *queryResolver) DepthSearch(ctx context.Context, input model.DepthFilter) ([]*model.Node, error) {
	if input.Reverse != nil && *input.Reverse {
		return r.store.DepthTo(ctx, input)
	}
	return r.store.DepthFrom(ctx, input)
}

func (r *queryResolver) GetEdge(ctx context.Context, input model.Path) (*model.Edge, error) {
	return r.store.Edge(ctx, input)
}

func (r *queryResolver) GetEdges(ctx context.Context, input model.Filter) ([]*model.Edge, error) {
	return r.store.Edges(ctx, input)
}

func (r *subscriptionResolver) NodeChange(ctx context.Context, typeArg model.ChangeFilter) (<-chan *model.Node, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	filter := func(obj interface{}) bool {
		node, ok := obj.(*model.Node)
		if !ok {
			return false
		}
		pass, _ := model.Evaluate(typeArg.Expressions, node)
		return pass
	}
	ch := make(chan *model.Node)
	r.machine.Go(func(routine machine.Routine) {
		routine.SubscribeFilter(typeArg.Type, filter, func(msg interface{}) {
			for {
				select {
				case <-routine.Context().Done():
					return
				case <-ctx.Done():
					return
				default:
					ch <- msg.(*model.Node)
				}
			}
		})
	})
	return ch, nil
}

func (r *subscriptionResolver) EdgeChange(ctx context.Context, typeArg model.ChangeFilter) (<-chan *model.Edge, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	filter := func(obj interface{}) bool {
		edge, ok := obj.(*model.Edge)
		if !ok {
			return false
		}
		pass, _ := model.Evaluate(typeArg.Expressions, edge)
		return pass
	}
	ch := make(chan *model.Edge)
	r.machine.Go(func(routine machine.Routine) {
		routine.SubscribeFilter(typeArg.Type, filter, func(msg interface{}) {
			for {
				select {
				case <-routine.Context().Done():
					return
				case <-ctx.Done():
					return
				default:
					ch <- msg.(*model.Edge)
				}
			}
		})
	})
	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
