package inspector

import (
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	astinspector "golang.org/x/tools/go/ast/inspector"

	kalerrors "github.com/JoelSpeed/kal/pkg/analysis/errors"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/extractjsontags"
	"github.com/JoelSpeed/kal/pkg/analysis/helpers/markers"
)

const name = "inspector"

// Analyzer is the analyzer for the inspector package.
// It provides common functionality for analyzers that need to inspect fields and struct.
// Abstracting away filtering of fields that the analyzers should and shouldn't be worrying about.
var Analyzer = &analysis.Analyzer{
	Name:       name,
	Doc:        "Provides common functionality for analyzers that need to inspect fields and struct",
	Run:        run,
	Requires:   []*analysis.Analyzer{inspect.Analyzer, extractjsontags.Analyzer, markers.Analyzer},
	ResultType: reflect.TypeOf(newInspector(nil, nil, nil)),
}

func run(pass *analysis.Pass) (any, error) {
	astinspector, ok := pass.ResultOf[inspect.Analyzer].(*astinspector.Inspector)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetInspector
	}

	jsonTags, ok := pass.ResultOf[extractjsontags.Analyzer].(extractjsontags.StructFieldTags)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetJSONTags
	}

	markersAccess, ok := pass.ResultOf[markers.Analyzer].(markers.Markers)
	if !ok {
		return nil, kalerrors.ErrCouldNotGetMarkers
	}

	return newInspector(astinspector, jsonTags, markersAccess), nil
}
