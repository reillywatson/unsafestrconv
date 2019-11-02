package unsafestrconv

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const Doc = `check for string(int) conversions`

var Analyzer = &analysis.Analyzer{
	Doc:      Doc,
	Name:     "unsafestrconv",
	Run:      unsafestrconvCheck,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func unsafestrconvCheck(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch fn := n.(type) {
			case *ast.CallExpr:
				switch fnName := fn.Fun.(type) {
				case *ast.Ident:
					if fnName.Name == "string" {
						if len(fn.Args) == 1 {
							switch arg := fn.Args[0].(type) {
							case *ast.Ident:
								obj := pass.TypesInfo.ObjectOf(arg)
								if obj != nil && obj.Type().Underlying().String() == "int" {
									reportNodef(pass, n, "Unsafe string(int) conversion")
								}
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func reportNodef(pass *analysis.Pass, node ast.Node, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	pass.Report(analysis.Diagnostic{Pos: node.Pos(), End: node.End(), Message: msg})
}
