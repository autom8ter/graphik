// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Counter struct {
	Count int `json:"count"`
}

type DepthFilter struct {
	Depth       int      `json:"depth"`
	Path        Path     `json:"path"`
	EdgeType    string   `json:"edgeType"`
	Reverse     *bool    `json:"reverse"`
	Expressions []string `json:"expressions"`
	Limit       int      `json:"limit"`
}

type EdgeConstructor struct {
	Path       Path                   `json:"path"`
	Mutual     bool                   `json:"mutual"`
	Attributes map[string]interface{} `json:"attributes"`
	From       Path                   `json:"from"`
	To         Path                   `json:"to"`
}

type Export struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

type Filter struct {
	Type        string   `json:"type"`
	Expressions []string `json:"expressions"`
	Limit       int      `json:"limit"`
}

type NodeConstructor struct {
	Path       Path                   `json:"path"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Patch struct {
	Path  Path                   `json:"path"`
	Patch map[string]interface{} `json:"patch"`
}

type Search struct {
	Search string `json:"search"`
	Type   string `json:"type"`
	Limit  int    `json:"limit"`
}

type SearchResult struct {
	Path Path        `json:"path"`
	Val  interface{} `json:"val"`
}

type SearchResults struct {
	Search  string          `json:"search"`
	Results []*SearchResult `json:"results"`
}

type Operator string

const (
	OperatorNeq Operator = "NEQ"
	OperatorEq  Operator = "EQ"
)

var AllOperator = []Operator{
	OperatorNeq,
	OperatorEq,
}

func (e Operator) IsValid() bool {
	switch e {
	case OperatorNeq, OperatorEq:
		return true
	}
	return false
}

func (e Operator) String() string {
	return string(e)
}

func (e *Operator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Operator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Operator", str)
	}
	return nil
}

func (e Operator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
