package parsy

import "go/ast"

type InspectorOptions struct {
	AvoidNil bool
}

type Inspector struct {
	ops            *InspectorOptions
	All            listeners[ast.Node]
	ArrayType      listeners[*ast.ArrayType]
	AssignStmt     listeners[*ast.AssignStmt]
	BadDecl        listeners[*ast.BadDecl]
	BadExpr        listeners[*ast.BadExpr]
	BadStmt        listeners[*ast.BadStmt]
	BasicLit       listeners[*ast.BasicLit]
	BinaryExpr     listeners[*ast.BinaryExpr]
	BlockStmt      listeners[*ast.BlockStmt]
	BranchStmt     listeners[*ast.BranchStmt]
	CallExpr       listeners[*ast.CallExpr]
	CaseClause     listeners[*ast.CaseClause]
	ChanType       listeners[*ast.ChanType]
	CommClause     listeners[*ast.CommClause]
	Comment        listeners[*ast.Comment]
	CommentGroup   listeners[*ast.CommentGroup]
	CompositeLit   listeners[*ast.CompositeLit]
	DeclStmt       listeners[*ast.DeclStmt]
	DeferStmt      listeners[*ast.DeferStmt]
	Ellipsis       listeners[*ast.Ellipsis]
	EmptyStmt      listeners[*ast.EmptyStmt]
	ExprStmt       listeners[*ast.ExprStmt]
	Field          listeners[*ast.Field]
	FieldList      listeners[*ast.FieldList]
	File           listeners[*ast.File]
	ForStmt        listeners[*ast.ForStmt]
	FuncDecl       listeners[*ast.FuncDecl]
	FuncLit        listeners[*ast.FuncLit]
	FuncType       listeners[*ast.FuncType]
	GenDecl        listeners[*ast.GenDecl]
	GoStmt         listeners[*ast.GoStmt]
	Ident          listeners[*ast.Ident]
	IfStmt         listeners[*ast.IfStmt]
	ImportSpec     listeners[*ast.ImportSpec]
	IncDecStmt     listeners[*ast.IncDecStmt]
	IndexExpr      listeners[*ast.IndexExpr]
	IndexListExpr  listeners[*ast.IndexListExpr]
	InterfaceType  listeners[*ast.InterfaceType]
	KeyValueExpr   listeners[*ast.KeyValueExpr]
	LabeledStmt    listeners[*ast.LabeledStmt]
	MapType        listeners[*ast.MapType]
	_package       listeners[*ast.Package]
	ParenExpr      listeners[*ast.ParenExpr]
	RangeStmt      listeners[*ast.RangeStmt]
	ReturnStmt     listeners[*ast.ReturnStmt]
	SelectStmt     listeners[*ast.SelectStmt]
	SelectorExpr   listeners[*ast.SelectorExpr]
	SendStmt       listeners[*ast.SendStmt]
	SliceExpr      listeners[*ast.SliceExpr]
	StarExpr       listeners[*ast.StarExpr]
	StructType     listeners[*ast.StructType]
	SwitchStmt     listeners[*ast.SwitchStmt]
	TypeAssertExpr listeners[*ast.TypeAssertExpr]
	TypeSpec       listeners[*ast.TypeSpec]
	TypeSwitchStmt listeners[*ast.TypeSwitchStmt]
	UnaryExpr      listeners[*ast.UnaryExpr]
	ValueSpec      listeners[*ast.ValueSpec]
}

func (i *Inspector) Visit(n ast.Node) ast.Visitor {
	if i.ops.AvoidNil && n == nil {
		return i
	}

	switch t := n.(type) {
	case *ast.ArrayType:
		i.ArrayType.trigger(t)
	case *ast.AssignStmt:
		i.AssignStmt.trigger(t)
	case *ast.BadDecl:
		i.BadDecl.trigger(t)
	case *ast.BadExpr:
		i.BadExpr.trigger(t)
	case *ast.BadStmt:
		i.BadStmt.trigger(t)
	case *ast.BasicLit:
		i.BasicLit.trigger(t)
	case *ast.BinaryExpr:
		i.BinaryExpr.trigger(t)
	case *ast.BlockStmt:
		i.BlockStmt.trigger(t)
	case *ast.BranchStmt:
		i.BranchStmt.trigger(t)
	case *ast.CallExpr:
		i.CallExpr.trigger(t)
	case *ast.CaseClause:
		i.CaseClause.trigger(t)
	case *ast.ChanType:
		i.ChanType.trigger(t)
	case *ast.CommClause:
		i.CommClause.trigger(t)
	case *ast.Comment:
		i.Comment.trigger(t)
	case *ast.CommentGroup:
		i.CommentGroup.trigger(t)
	case *ast.CompositeLit:
		i.CompositeLit.trigger(t)
	case *ast.DeclStmt:
		i.DeclStmt.trigger(t)
	case *ast.DeferStmt:
		i.DeferStmt.trigger(t)
	case *ast.Ellipsis:
		i.Ellipsis.trigger(t)
	case *ast.EmptyStmt:
		i.EmptyStmt.trigger(t)
	case *ast.ExprStmt:
		i.ExprStmt.trigger(t)
	case *ast.Field:
		i.Field.trigger(t)
	case *ast.FieldList:
		i.FieldList.trigger(t)
	case *ast.File:
		i.File.trigger(t)
	case *ast.ForStmt:
		i.ForStmt.trigger(t)
	case *ast.FuncDecl:
		i.FuncDecl.trigger(t)
	case *ast.FuncLit:
		i.FuncLit.trigger(t)
	case *ast.FuncType:
		i.FuncType.trigger(t)
	case *ast.GenDecl:
		i.GenDecl.trigger(t)
	case *ast.GoStmt:
		i.GoStmt.trigger(t)
	case *ast.Ident:
		i.Ident.trigger(t)
	case *ast.IfStmt:
		i.IfStmt.trigger(t)
	case *ast.ImportSpec:
		i.ImportSpec.trigger(t)
	case *ast.IncDecStmt:
		i.IncDecStmt.trigger(t)
	case *ast.IndexExpr:
		i.IndexExpr.trigger(t)
	case *ast.IndexListExpr:
		i.IndexListExpr.trigger(t)
	case *ast.InterfaceType:
		i.InterfaceType.trigger(t)
	case *ast.KeyValueExpr:
		i.KeyValueExpr.trigger(t)
	case *ast.LabeledStmt:
		i.LabeledStmt.trigger(t)
	case *ast.MapType:
		i.MapType.trigger(t)
	case *ast.Package:
		i._package.trigger(t)
	case *ast.ParenExpr:
		i.ParenExpr.trigger(t)
	case *ast.RangeStmt:
		i.RangeStmt.trigger(t)
	case *ast.ReturnStmt:
		i.ReturnStmt.trigger(t)
	case *ast.SelectStmt:
		i.SelectStmt.trigger(t)
	case *ast.SelectorExpr:
		i.SelectorExpr.trigger(t)
	case *ast.SendStmt:
		i.SendStmt.trigger(t)
	case *ast.SliceExpr:
		i.SliceExpr.trigger(t)
	case *ast.StarExpr:
		i.StarExpr.trigger(t)
	case *ast.StructType:
		i.StructType.trigger(t)
	case *ast.SwitchStmt:
		i.SwitchStmt.trigger(t)
	case *ast.TypeAssertExpr:
		i.TypeAssertExpr.trigger(t)
	case *ast.TypeSpec:
		i.TypeSpec.trigger(t)
	case *ast.TypeSwitchStmt:
		i.TypeSwitchStmt.trigger(t)
	case *ast.UnaryExpr:
		i.UnaryExpr.trigger(t)
	case *ast.ValueSpec:
		i.ValueSpec.trigger(t)
	}
	i.All.trigger(n)
	return i
}

// NewInspector creates a new Inspector pointer with
// the given options
func NewInspector(ops *InspectorOptions) *Inspector {
	if ops == nil {
		ops = &InspectorOptions{
			AvoidNil: false,
		}
	}

	return &Inspector{
		ops: ops,
	}
}

// InspectNode inspects a node recursive
func InspectNode(i *Inspector, n ast.Node) {
	ast.Walk(i, n)
}
