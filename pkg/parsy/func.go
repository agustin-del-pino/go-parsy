package parsy

import "go/ast"

type ReturnCaseType int

type FuncRecv struct {
	IsPointer bool
	Name      string
	Ident     *ast.Ident
}

type CastExpr struct {
	Expr ast.Expr
}

func (c *CastExpr) ToBasicLit() (*ast.BasicLit, bool) {
	return AsBasicLit(c.Expr)
}

func (c *CastExpr) ToIdent() (*ast.Ident, bool) {
	return AsIdent(c.Expr)
}

func (c *CastExpr) ToStarExpr() (*ast.StarExpr, bool) {
	return AsStarExpr(c.Expr)
}

func (c *CastExpr) ToUnaryExpr() (*ast.UnaryExpr, bool) {
	return AsUnaryExpr(c.Expr)
}

func (c *CastExpr) ToCompositeLit() (*ast.CompositeLit, bool) {
	return AsCompositeLit(c.Expr)
}

type FuncReturnType struct {
	IsPointer bool
	Ident     *ast.Ident
}

type FuncReturnCase struct {
	Results []*CastExpr
}

type FuncReturns struct {
	Types       []*FuncReturnType
	ReturnCases []*FuncReturnCase
}

type FuncInfo struct {
	Name      string
	Receivers []*FuncRecv
	Returns   *FuncReturns
}

func GetFuncReceivers(fn *ast.FuncDecl) ([]*FuncRecv, bool) {
	if fn.Recv == nil {
		return nil, false
	}

	frs := make([]*FuncRecv, len(fn.Recv.List))

	for i := range frs {
		frs[i] = &FuncRecv{
			Name: fn.Recv.List[i].Names[0].Name,
		}

		var rcvT ast.Node

		if se, ok := AsStarExpr(fn.Recv.List[i].Type); ok {
			frs[i].IsPointer = true
			rcvT = se.X
		} else {
			rcvT = fn.Recv.List[i].Type
		}

		if idt, ok := AsIdent(rcvT); !ok {
			return nil, false
		} else {
			frs[i].Ident = idt
		}
	}

	return frs, true
}

func GetFuncName(fn *ast.FuncDecl) (string, bool) {
	if fn.Name == nil {
		return "", false
	}
	return fn.Name.Name, true
}

func GetFuncReturnTypes(fn *ast.FuncDecl) ([]*FuncReturnType, bool) {
	if fn.Type == nil || fn.Type.Results == nil {
		return nil, false
	}

	ft := make([]*FuncReturnType, len(fn.Type.Results.List))

	for i := range ft {
		ft[i] = &FuncReturnType{}

		n := fn.Type.Results.List[i].Type

		if se, ok := AsStarExpr(n); ok {
			ft[i].IsPointer = true
			n = se.X
		}

		if idt, ok := AsIdent(n); ok {
			ft[i].Ident = idt
		} else {
			return nil, false
		}
	}

	return ft, true
}

func GetFuncReturnCases(fn *ast.FuncDecl) ([]*FuncReturnCase, bool) {
	i := NewInspector(&InspectorOptions{
		AvoidNil: true,
	})

	var (
		frc []*FuncReturnCase
		ok  bool
	)

	i.ReturnStmt.AddListener(func(rs *ast.ReturnStmt) {
		ok = true
		fc := &FuncReturnCase{
			Results: make([]*CastExpr, len(rs.Results)),
		}

		for i := range fc.Results {
			fc.Results[i] = &CastExpr{
				Expr: rs.Results[i],
			}
		}
		frc = append(frc, fc)
	})

	InspectNode(i, fn)

	return frc, ok
}

func GetFuncInfo(fn *ast.FuncDecl) *FuncInfo {
	fi := &FuncInfo{}

	if n, ok := GetFuncName(fn); ok {
		fi.Name = n
	}

	if r, ok := GetFuncReceivers(fn); ok {
		fi.Receivers = r
	}

	if rt, ok := GetFuncReturnTypes(fn); ok {
		if fi.Returns == nil {
			fi.Returns = &FuncReturns{}
		}
		fi.Returns.Types = rt
	}

	if rc, ok := GetFuncReturnCases(fn); ok {
		if fi.Returns == nil {
			fi.Returns = &FuncReturns{}
		}
		fi.Returns.ReturnCases = rc
	}

	return fi
}
