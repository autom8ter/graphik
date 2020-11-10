package lang

import (
	"github.com/autom8ter/graphik/graph"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/interpreter/functions"
	"github.com/pkg/errors"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

var dynamic = decls.NewMapType(decls.String, decls.Any)

var defaultFields = []*exprpb.Decl{
	decls.NewVar("input", dynamic),
}

type Function struct {
	Overload    *functions.Overload
	Declaration *Declaration
}

type Declaration struct {
	Name      string
	Overloads []*exprpb.Decl_FunctionDecl_Overload
}

type FuncMap map[string]*Function

func (f FuncMap) functions() []*functions.Overload {
	var fns []*functions.Overload
	for _, fn := range f {
		fns = append(fns, fn.Overload)
	}
	return fns
}

func (f FuncMap) MapEval(expression string, args map[string]interface{}) (map[string]interface{}, *cel.EvalDetails, error) {
	var declarations []*exprpb.Decl
	for name, fn := range f {
		declarations = append(declarations, decls.NewFunction(name, fn.Declaration.Overloads...))
	}
	for _, ext := range defaultFields {
		declarations = append(declarations, ext)
	}
	env, err := cel.NewEnv(cel.Declarations(declarations...))
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating env")
	}
	ast, iss := env.Compile(expression)
	if iss.Err() != nil {
		return nil, nil, errors.Wrap(iss.Err(), "compiling env")
	}

	program, err := env.Program(ast, cel.Functions(f.functions()...))
	if err != nil {
		return nil, nil, errors.Wrap(err, "programming env")
	}
	result, details, err := program.Eval(map[string]interface{}{
		"input": args,
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "evaluating program")
	}
	return graph.ToMap(result), details, nil
}

func (f FuncMap) BoolEval(expressions []string, args map[string]interface{}) (bool, error) {
	if len(expressions) == 0 {
		return true, nil
	}
	var declarations []*exprpb.Decl
	for name, fn := range f {
		declarations = append(declarations, decls.NewFunction(name, fn.Declaration.Overloads...))
	}
	for _, ext := range defaultFields {
		declarations = append(declarations, ext)
	}
	env, err := cel.NewEnv(cel.Declarations(declarations...))
	if err != nil {
		return false, errors.Wrap(err, "creating env")
	}
	var programs []cel.Program
	for _, exp := range expressions {
		ast, iss := env.Compile(exp)
		if iss.Err() != nil {
			return false, iss.Err()
		}
		program, err := env.Program(ast, cel.Functions(f.functions()...))
		if err != nil {
			return false, err
		}
		programs = append(programs, program)
	}
	var passes = true
	for _, program := range programs {
		out, _, err := program.Eval(map[string]interface{}{
			"input": args,
		})
		if err != nil {
			return false, err
		}
		if val, ok := out.Value().(bool); !ok || !val {
			passes = false
		}
	}
	return passes, nil
}