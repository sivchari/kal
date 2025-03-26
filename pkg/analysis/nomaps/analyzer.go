package nomaps

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"github.com/JoelSpeed/kal/pkg/analysis/helpers/extractjsontags"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/inspector"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/markers"
	"github.com/JoelSpeed/kal/pkg/config"
	"golang.org/x/tools/go/analysis"
)

const (
	name = "nomaps"
)

var (
	errCouldNotGetInspector = errors.New("could not get inspector")
)

type analyzer struct {
	policy config.NoMapsPolicy
}

// newAnalyzer creates a new analyzer.
func newAnalyzer(cfg config.NoMapsConfig) *analysis.Analyzer {
	defaultConfig(&cfg)

	a := &analyzer{
		policy: cfg.Policy,
	}

	return &analysis.Analyzer{
		Name:     name,
		Doc:      "Checks for usage of map types. Maps are discouraged apart from `map[string]string` which is used for labels and annotations. Use a list of named objects instead.",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspector.Analyzer},
	}
}

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspect, ok := pass.ResultOf[inspector.Analyzer].(inspector.Inspector)
	if !ok {
		return nil, errCouldNotGetInspector
	}

	inspect.InspectFields(func(field *ast.Field, stack []ast.Node, jsonTagInfo extractjsontags.FieldTagInfo, markersAccess markers.Markers) {
		a.checkField(pass, field)
	})

	return nil, nil //nolint:nilnil
}

func (a *analyzer) checkField(pass *analysis.Pass, field *ast.Field) {
	stringToStringMapType := types.NewMap(types.Typ[types.String], types.Typ[types.String])

	underlyingType := pass.TypesInfo.TypeOf(field.Type).Underlying()

	if ptr, ok := underlyingType.(*types.Pointer); ok {
		underlyingType = ptr.Elem().Underlying()
	}

	m, ok := underlyingType.(*types.Map)
	if !ok {
		return
	}

	if a.policy == config.NoMapsEnforce {
		report(pass, field.Pos(), field.Names[0].Name)
		return
	}

	if a.policy == config.NoMapsAllowStringToStringMaps {
		if types.Identical(m, stringToStringMapType) {
			return
		}

		report(pass, field.Pos(), field.Names[0].Name)
	}

	if a.policy == config.NoMapsIgnore {
		key := m.Key().Underlying()
		_, ok := key.(*types.Basic)

		elm := m.Elem().Underlying()
		_, ok2 := elm.(*types.Basic)

		if ok && ok2 {
			return
		}

		report(pass, field.Pos(), field.Names[0].Name)
	}
}

func report(pass *analysis.Pass, pos token.Pos, fieldName string) {
	pass.Report(analysis.Diagnostic{
		Pos:     pos,
		Message: fmt.Sprintf("%s should not use a map type, use a list type with a unique name/identifier instead", fieldName),
	})
}

func defaultConfig(cfg *config.NoMapsConfig) {
	if cfg.Policy == "" {
		cfg.Policy = config.NoMapsAllowStringToStringMaps
	}
}
