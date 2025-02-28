package jsontags

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	"github.com/JoelSpeed/kal/pkg/analysis/helpers/extractjsontags"
	"github.com/JoelSpeed/kal/pkg/config"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	// camelCaseRegex is a regular expression that matches camel case strings.
	camelCaseRegex = "^[a-z][a-z0-9]*(?:[A-Z][a-z0-9]*)*$"

	name = "jsontags"
)

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
	errCouldNotGetJSONTags  = errors.New("could not get json tags")
)

type analyzer struct {
	jsonTagRegex *regexp.Regexp
}

// newAnalyzer creates a new analyzer with the given json tag regex.
func newAnalyzer(cfg config.JSONTagsConfig) (*analysis.Analyzer, error) {
	defaultConfig(&cfg)

	jsonTagRegex, err := regexp.Compile(cfg.JSONTagRegex)
	if err != nil {
		return nil, fmt.Errorf("could not compile json tag regex: %w", err)
	}

	a := &analyzer{
		jsonTagRegex: jsonTagRegex,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Check that all struct fields in an API are tagged with json tags",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer},
	}, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	inspect, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	jsonTags, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		return nil, errCouldNotGetJSONTags
	}

	// Filter to fields so that we can iterate over fields in a struct.
	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) (proceed bool) {
		if !push {
			return false
		}

		if len(stack) < 2 {
			return true
		}

		// The 0th node in the stack is the *ast.File.
		// The 1st node in the stack is the *ast.GenDecl.
		decl, ok := stack[1].(*ast.GenDecl)
		if !ok {
			return false
		}

		if decl.Tok != token.TYPE {
			// Returning false here means we won't inspect non-type declarations (e.g. var, const, import).
			return false
		}

		field, ok := n.(*ast.Field)
		if !ok {
			return true
		}

		return a.checkField(pass, field, jsonTags)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field, jsonTags extractjsontags.StructFieldTags) (proceed bool) {
	tagInfo := jsonTags.FieldTags(field)

	var prefix string
	if len(field.Names) > 0 && field.Names[0] != nil {
		prefix = fmt.Sprintf("field %s", field.Names[0].Name)
	} else if ident, ok := field.Type.(*ast.Ident); ok {
		prefix = fmt.Sprintf("embedded field %s", ident.Name)
	}

	if tagInfo.Missing {
		pass.Reportf(field.Pos(), "%s is missing json tag", prefix)
		return true
	}

	if tagInfo.Inline {
		return true
	}

	if tagInfo.Ignored {
		// Returning false here means we won't inspect the children of an ignored field.
		return false
	}

	if tagInfo.Name == "" {
		pass.Reportf(field.Pos(), "%s has empty json tag", prefix)
		return true
	}

	matched := a.jsonTagRegex.Match([]byte(tagInfo.Name))
	if !matched {
		pass.Reportf(field.Pos(), "%s json tag does not match pattern %q: %s", prefix, a.jsonTagRegex.String(), tagInfo.Name)
	}

	return true
}

func defaultConfig(cfg *config.JSONTagsConfig) {
	if cfg.JSONTagRegex == "" {
		cfg.JSONTagRegex = camelCaseRegex
	}
}
