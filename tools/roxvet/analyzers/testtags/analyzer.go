package testtags

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/stackrox/rox/tools/roxvet/common"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `Ensure that our *_test.go files include a //go:build <TAG>, otherwise they won't be run.'`

const roxPrefix = "github.com/stackrox/rox/"

// Analyzer is a analysis.Analyzer from the analysis package of the Go standard lib. [It analyzes code]
var Analyzer = &analysis.Analyzer{
	Name:     "testtags",
	Doc:      doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Make sure this runs only on tests within stackrox/tests, our go e2e tests
func isTestsPackage(packageName string) bool {
	if !strings.HasPrefix(packageName, roxPrefix) {
		return false
	}
	unqualifiedPackageName := strings.TrimPrefix(packageName, roxPrefix)
	pathElems := strings.Split(unqualifiedPackageName, string(filepath.Separator))
	if len(pathElems) == 0 {
		return false
	}
	if pathElems[0] == "tests" {
		return true
	}
	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	if !isTestsPackage(pass.Pkg.Path()) {
		return nil, nil
	}

	common.FilteredPreorder(inspectResult, common.IsTestFile, nodeFilter, func(n ast.Node) {
		goBuildDirectiveCount := 0
		pos := token.NoPos
		fileNode := n.(*ast.File)
		pos = fileNode.Pos()
		for _, comment := range fileNode.Comments {
			if strings.HasPrefix(comment.Text(), "//go:build") {
				goBuildDirectiveCount++
			}
		}
		if goBuildDirectiveCount == 0 {
			pass.Reportf(pos, "Missing //go:build directive.")
		} else if goBuildDirectiveCount > 1 {
			pass.Reportf(pos, "Multiple //go:build directives, there should be exactly one.")
		}
	})
	return nil, nil
}
