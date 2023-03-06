package parsy

import "go/ast"

// As casts an ast.Node to a specific Node Type.
func As[T ast.Node](n ast.Node) (T, bool) {
	v, ok := n.(T)
	return v, ok
}

// AsIdent casts an ast.Node to *ast.Ident
func AsIdent(n ast.Node) (*ast.Ident, bool) {
	return As[*ast.Ident](n)
}

// AsStarExpr casts an ast.Node to *ast.StarExpr
func AsStarExpr(n ast.Node) (*ast.StarExpr, bool) {
	return As[*ast.StarExpr](n)
}

// AsBasicLit casts an ast.Node to *ast.BasicLit
func AsBasicLit(n ast.Node) (*ast.BasicLit, bool) {
	return As[*ast.BasicLit](n)
}

// AsUnaryExpr casts an ast.Node to *ast.UnaryExpr
func AsUnaryExpr(n ast.Node) (*ast.UnaryExpr, bool) {
	return As[*ast.UnaryExpr](n)
}

// AsCompositeLit casts an ast.Node to *ast.CompositeLit
func AsCompositeLit(n ast.Node) (*ast.CompositeLit, bool) {
	return As[*ast.CompositeLit](n)
}
