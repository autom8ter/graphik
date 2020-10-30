package generic

import (
	"github.com/autom8ter/graphik/graph/model"
	"github.com/jmespath/go-jmespath"
	"time"
)

type Nodes struct {
	nodes map[string]map[string]*model.Node
}

func NewNodes() *Nodes {
	return &Nodes{
		nodes: map[string]map[string]*model.Node{},
	}
}

func (n *Nodes) Len(nodeType string) int {
	if c, ok := n.nodes[nodeType]; ok {
		return len(c)
	}
	return 0
}

func (n *Nodes) Types() []string {
	var nodeTypes []string
	for k, _ := range n.nodes {
		nodeTypes = append(nodeTypes, k)
	}
	return nodeTypes
}

func (n *Nodes) All() []*model.Node {
	var nodes []*model.Node
	n.Range(Any, func(node *model.Node) bool {
		nodes = append(nodes, node)
		return true
	})
	return nodes
}

func (n *Nodes) Get(key model.ForeignKey) (*model.Node, bool) {
	if c, ok := n.nodes[key.Type]; ok {
		node := c[key.ID]
		return c[key.ID], node != nil
	}
	return nil, false
}

func (n *Nodes) Set(value *model.Node) *model.Node {
	if value.ID == "" {
		value.ID = uuid()
	}
	if _, ok := n.nodes[value.Type]; !ok {
		n.nodes[value.Type] = map[string]*model.Node{}
	}
	n.nodes[value.Type][value.ID] = value
	return value
}

func (n *Nodes) Patch(updatedAt time.Time, value *model.Patch) *model.Node {
	if _, ok := n.nodes[value.Type]; !ok {
		return nil
	}
	for k, v := range value.Patch {
		n.nodes[value.Type][value.ID].Attributes[k] = v
	}
	n.nodes[value.Type][value.ID].UpdatedAt = &updatedAt
	return n.nodes[value.Type][value.ID]
}

func (n *Nodes) Range(nodeType string, f func(node *model.Node) bool) {
	if nodeType == Any {
		for _, c := range n.nodes {
			for _, v := range c {
				f(v)
			}
		}
	} else {
		if c, ok := n.nodes[nodeType]; ok {
			for _, v := range c {
				f(v)
			}
		}
	}
}

func (n *Nodes) Delete(key model.ForeignKey) {
	if c, ok := n.nodes[key.Type]; ok {
		delete(c, key.ID)
	}
}

func (n *Nodes) Exists(key model.ForeignKey) bool {
	_, ok := n.Get(key)
	return ok
}

func (n *Nodes) Filter(nodeType string, filter func(node *model.Node) bool) []*model.Node {
	var filtered []*model.Node
	n.Range(nodeType, func(node *model.Node) bool {
		if filter(node) {
			filtered = append(filtered, node)
		}
		return true
	})
	return filtered
}

func (n *Nodes) SetAll(nodes ...*model.Node) []*model.Node {
	var created []*model.Node
	for _, node := range nodes {
		created = append(created, n.Set(node))
	}
	return created
}

func (n *Nodes) DeleteAll(Nodes ...*model.Node) {
	for _, node := range Nodes {
		n.Delete(model.ForeignKey{
			ID:   node.ID,
			Type: node.Type,
		})
	}
}

func (n *Nodes) Clear(nodeType string) {
	if cache, ok := n.nodes[nodeType]; ok {
		for k, _ := range cache {
			delete(cache, k)
		}
	}
}

func (n *Nodes) Close() {
	for nodeType, _ := range n.nodes {
		n.Clear(nodeType)
	}
}

func (n *Nodes) Search(expression, nodeType string) (*model.SearchResults, error) {
	results := &model.SearchResults{
		Search: expression,
	}
	exp, err := jmespath.Compile(expression)
	if err != nil {
		return nil, err
	}
	n.Range(nodeType, func(node *model.Node) bool {
		val, _ := exp.Search(node)
		if val != nil {
			results.Results = append(results.Results, &model.SearchResult{
				ID:   node.ID,
				Type: node.Type,
				Val:  val,
			})
		}
		return true
	})
	return results, nil
}


func (n *Nodes) FilterSearch(filter model.Filter) ([]*model.Node, error) {
	var nodes []*model.Node
	n.Range(filter.Type, func(node *model.Node) bool {
		for _, exp := range filter.Expressions {
			val, _ := jmespath.Search(exp.Key, node)
			if exp.Operator == model.OperatorNeq {
				if val == exp.Value {
					return true
				}
			}
			if exp.Operator == model.OperatorEq {
				if val != exp.Value {
					return true
				}
			}
		}
		nodes = append(nodes, node)
		return len(nodes) < filter.Limit
	})
	return nodes, nil
}
