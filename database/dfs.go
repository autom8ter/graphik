package database

import (
	"context"
	apipb "github.com/autom8ter/graphik/gen/go"
	"github.com/autom8ter/graphik/generic/stack"
	"github.com/autom8ter/graphik/logger"
	"github.com/google/cel-go/cel"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"strings"
)

// depthFirst implements stateful depth-first graph traversal.
type depthFirst struct {
	g       *Graph
	stack   *stack.Stack
	visited map[string]struct{}
	docs    *apipb.Traversals
	filter  *apipb.TFilter
}

func (g *Graph) newDepthFirst(filter *apipb.TFilter) *depthFirst {
	return &depthFirst{
		g:       g,
		filter:  filter,
		stack:   stack.New(),
		visited: map[string]struct{}{},
		docs:    &apipb.Traversals{},
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
			d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
				Doc:         doc,
				RelativeRef: &apipb.Refs{Refs: []*apipb.Ref{doc.GetRef()}},
			})
		} else {
			res, err := d.g.vm.Doc().Eval(doc, *docProgram)
			if err != nil {
				return err
			}
			if res {
				d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
					Doc:         doc,
					RelativeRef: &apipb.Refs{Refs: []*apipb.Ref{doc.GetRef()}},
				})
			}
		}
		d.visited[d.filter.Root.String()] = struct{}{}
	}
	var traversalRef []*apipb.Ref
	for d.stack.Len() > 0 && len(d.docs.Traversals) < int(d.filter.Limit) {
		if err := ctx.Err(); err != nil {
			return nil
		}
		popped := d.stack.Pop().(*apipb.Doc)
		traversalRef = append(traversalRef, popped.GetRef())
		if len(d.docs.Traversals) >= int(d.filter.Limit) {
			return nil
		}
		if err := d.g.rangeFrom(ctx, tx, popped.GetRef(), func(e *apipb.Connection) bool {
			if connectionProgram != nil {
				res, err := d.g.vm.Connection().Eval(e, *connectionProgram)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if !res {
					return true
				}
			}
			if _, ok := d.visited[e.From.String()]; !ok {
				from, err := d.g.getDoc(ctx, tx, e.From)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if docProgram == nil {
					traversalRef = append(traversalRef, from.GetRef())
					d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
						Doc:         from,
						RelativeRef: &apipb.Refs{Refs: traversalRef},
						Direction:   apipb.Direction_From,
					})
				} else {
					res, err := d.g.vm.Doc().Eval(from, *docProgram)
					if err != nil {
						if !strings.Contains(err.Error(), "no such key") {
							logger.Error("dfs failure", zap.Error(err))
						}
						return true
					}
					if res {
						traversalRef = append(traversalRef, from.GetRef())
						d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
							Doc:         from,
							RelativeRef: &apipb.Refs{Refs: traversalRef},
							Direction:   apipb.Direction_From,
						})
					}
				}
				d.visited[from.Ref.String()] = struct{}{}
				d.stack.Push(from)
			}
			if _, ok := d.visited[e.To.String()]; !ok {
				to, err := d.g.getDoc(ctx, tx, e.To)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if docProgram == nil {
					traversalRef = append(traversalRef, to.GetRef())
					d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
						Doc:         to,
						RelativeRef: &apipb.Refs{Refs: traversalRef},
						Direction:   apipb.Direction_From,
					})
				} else {
					res, err := d.g.vm.Doc().Eval(to, *docProgram)
					if err != nil {
						if !strings.Contains(err.Error(), "no such key") {
							logger.Error("dfs failure", zap.Error(err))
						}
						return true
					}
					if res {
						traversalRef = append(traversalRef, to.GetRef())
						d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
							Doc:         to,
							RelativeRef: &apipb.Refs{Refs: traversalRef},
							Direction:   apipb.Direction_From,
						})
					}
				}
				d.visited[to.Ref.String()] = struct{}{}
				d.stack.Push(to)
			}
			return len(d.docs.Traversals) < int(d.filter.Limit)
		}); err != nil {
			return err
		}
		if err := d.g.rangeTo(ctx, tx, popped.GetRef(), func(e *apipb.Connection) bool {
			if connectionProgram != nil {
				res, err := d.g.vm.Connection().Eval(e, *connectionProgram)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if !res {
					return true
				}
			}
			if _, ok := d.visited[e.From.String()]; !ok {
				from, err := d.g.getDoc(ctx, tx, e.From)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if docProgram == nil {
					traversalRef = append(traversalRef, from.GetRef())
					d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
						Doc:         from,
						RelativeRef: &apipb.Refs{Refs: traversalRef},
						Direction:   apipb.Direction_To,
					})
				} else {
					res, err := d.g.vm.Doc().Eval(from, *docProgram)
					if err != nil {
						if !strings.Contains(err.Error(), "no such key") {
							logger.Error("dfs failure", zap.Error(err))
						}
						return true
					}
					if res {
						traversalRef = append(traversalRef, from.GetRef())
						d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
							Doc:         from,
							RelativeRef: &apipb.Refs{Refs: traversalRef},
							Direction:   apipb.Direction_To,
						})
					}
				}
				d.visited[from.Ref.String()] = struct{}{}
				d.stack.Push(from)
			}
			if _, ok := d.visited[e.To.String()]; !ok {
				to, err := d.g.getDoc(ctx, tx, e.To)
				if err != nil {
					if !strings.Contains(err.Error(), "no such key") {
						logger.Error("dfs failure", zap.Error(err))
					}
					return true
				}
				if docProgram == nil {
					traversalRef = append(traversalRef, to.GetRef())
					d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
						Doc:         to,
						RelativeRef: &apipb.Refs{Refs: traversalRef},
						Direction:   apipb.Direction_To,
					})
				} else {
					res, err := d.g.vm.Doc().Eval(to, *docProgram)
					if err != nil {
						if !strings.Contains(err.Error(), "no such key") {
							logger.Error("dfs failure", zap.Error(err))
						}
						return true
					}
					if res {
						traversalRef = append(traversalRef, to.GetRef())
						d.docs.Traversals = append(d.docs.Traversals, &apipb.Traversal{
							Doc:         to,
							RelativeRef: &apipb.Refs{Refs: traversalRef},
							Direction:   apipb.Direction_To,
						})
					}
				}
				d.visited[to.Ref.String()] = struct{}{}
				d.stack.Push(to)
			}
			return len(d.docs.Traversals) < int(d.filter.Limit)
		}); err != nil {
			return err
		}
	}
	return nil
}
