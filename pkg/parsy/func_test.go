package parsy

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

const src = `
package main;

func Foo(i int) (string, error) {
	return string(i), nil
}

`

func TestGetFuncInfo(t *testing.T) {
	fn := setUp()
	
	fi := GetFuncInfo(fn)

	assert.NotNil(t, fi)
	assert.Equal(t, "Foo", fi.Name)
	assert.Nil(t, fi.Receivers)
	assert.NotNil(t, fi.Returns)
	
	assert.Len(t, fi.Returns.Types, 2)
	assert.False(t, fi.Returns.Types[0].IsPointer)
	assert.NotNil(t, fi.Returns.Types[0].Ident)
	assert.Equal(t, "string", fi.Returns.Types[0].Ident.Name)
	assert.False(t, fi.Returns.Types[1].IsPointer)
	assert.NotNil(t, fi.Returns.Types[1].Ident)
	assert.Equal(t, "error", fi.Returns.Types[1].Ident.Name)

	assert.Len(t, fi.Returns.ReturnCases, 1)
	assert.NotNil(t, fi.Returns.ReturnCases[0].Results)
	assert.Len(t, fi.Returns.ReturnCases[0].Results, 2)
	assert.NotNil(t, fi.Returns.ReturnCases[0].Results[0])
	assert.NotNil(t, fi.Returns.ReturnCases[0].Results[1])
}

func setUp() *ast.FuncDecl {
	i := NewInspector(&InspectorOptions{AvoidNil: true})
	var fn *ast.FuncDecl

	i.FuncDecl.AddListener(func(fd *ast.FuncDecl) {
		fn = fd
	})

	ParseString(i, src)

	return fn
}
