package checks

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const CaseCheck = "case"

type CaseAnalyzer struct {
	pass *analysis.Pass
}

func NewCaseAnalyzer(pass *analysis.Pass) *CaseAnalyzer {
	return &CaseAnalyzer{
		pass: pass,
	}
}

func (a *CaseAnalyzer) NodeFilter() []ast.Node {
	return []ast.Node{
		(*ast.CaseClause)(nil),
	}
}

func (a *CaseAnalyzer) Check(n ast.Node) {
	caseClause, ok := n.(*ast.CaseClause)
	if !ok {
		return
	}

	for _, c := range caseClause.List {
		switch x := c.(type) {
		case *ast.BasicLit:
			if isMagicNumber(x) {
				a.pass.Reportf(x.Pos(), reportMsg, x.Value, CaseCheck)
			}
		}
	}
}
