package database

import (
	"context"
	apipb "github.com/autom8ter/graphik/gen/go"
	"github.com/autom8ter/graphik/logger"
	"github.com/autom8ter/graphik/stack"
	"github.com/google/cel-go/cel"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

// depthFirst implements stateful depth-first graph traversal.
type depthFirst struct {
	g       *Graph
	stack   *stack.Stack
	visited map[string]struct{}
	docs    *apipb.Docs
	filter  *apipb.DepthFilter
}

func (g *Graph) NewdepthFirst(filter *apipb.DepthFilter) *depthFirst {
	return &depthFirst{
		g:       g,
		filter:  filter,
		stack:   stack.New(),
		visited: map[string]struct{}{},
		docs:    &apipb.Docs{},
	}
}

func (d *depthFirst) Walk(ctx context.Context, tx *bbolt.Tx) error {
	defer func() {
		d.docs.Sort(d.filter.Sort)
	}()
	var (
		docProgram        *cel.Program
		connectionProgram *cel.Program
	)
	if d.filter.GetConnectionExpression() != "" {
		program, err := d.g.vm.Connection().Program(d.filter.GetConnectionExpression())
		if err != nil {
			return err
		}
		connectionProgram = &program
	}
	if d.filter.GetDocExpression() != "" {
		program, err := d.g.vm.Doc().Program(d.filter.GetDocExpression())
		if err != nil {
			return err
		}
		docProgram = &program
	}
	if _, ok := d.visited[d.filter.Root.String()]; !ok {
		doc, err := d.g.getDoc(ctx, tx, d.filter.Root)
		if err != nil {
			return err
		}
		d.stack.Push(doc)
		if docProgram == nil {
			d.docs.Docs = append(d.docs.Docs, doc)
		} else {
			res, err := d.g.vm.Doc().Eval(doc, *docProgram)
			if err != nil {
				return err
			}
			if res {
				d.docs.Docs = append(d.docs.Docs, doc)
			}
		}
		d.visited[d.filter.Root.String()] = struct{}{}
	}
	for d.stack.Len() > 0 {
		t := d.stack.Pop()
		if len(d.docs.Docs) >= int(d.filter.Limit) {
			return nil
		}
		if err := d.g.rangeFrom(ctx, tx, t.(*apipb.Doc).GetPath(), func(e *apipb.Connection) bool {
			if connectionProgram != nil {
				res, err := d.g.vm.Connection().Eval(e, *connectionProgram)
				if err != nil {
					logger.Error("dfs failure", zap.Error(err))
					return true
				}
				if !res {
					return true
				}
			}
			if _, ok := d.visited[e.To.String()]; !ok {
				to, err := d.g.getDoc(ctx, tx, e.To)
				if err != nil {
					logger.Error("dfs failure", zap.Error(err))
					return true
				}
				if docProgram == nil {
					d.docs.Docs = append(d.docs.Docs, to)
				} else {
					res, err := d.g.vm.Doc().Eval(to, *docProgram)
					if err != nil {
						logger.Error("dfs failure", zap.Error(err))
						return true
					}
					if res {
						d.docs.Docs = append(d.docs.Docs, to)
					}
				}
				d.visited[to.String()] = struct{}{}
				d.stack.Push(to)
			}
			return len(d.docs.Docs) >= int(d.filter.Limit)
		}); err != nil {
			return err
		}
	}
	return nil
}
