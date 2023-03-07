package parsy

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFuncInfo_with_no_receiver(t *testing.T) {
	const src = `
package main;

func Foo(i int) (string, error) {
	return string(i), nil
}

`
	fn := setUp(src)

	fi := GetFuncInfo(fn)
	commonAsserts(t, fi)
	assert.Nil(t, fi.Receivers)

}

func TestGetFuncInfo_with_receiver(t *testing.T) {
	const src = `
	package main

	type MyStruct struct {}

	func (m MyStruct) Foo() (string, error) {
		return "hello", nil
	}

	`
	fn := setUp(src)

	fi := GetFuncInfo(fn)

	commonAsserts(t, fi)
	commonReceiverAsserts(t, fi)
	assert.False(t, fi.Receivers[0].IsPointer)
}

func TestGetFuncInfo_with_star_receiver(t *testing.T) {
	const src = `
	package main

	type MyStruct struct {}

	func (m *MyStruct) Foo() (string, error) {
		return "hello", nil
	}

	`
	fn := setUp(src)

	fi := GetFuncInfo(fn)

	commonAsserts(t, fi)
	commonReceiverAsserts(t, fi)
	assert.True(t, fi.Receivers[0].IsPointer)
}

func commonAsserts(t *testing.T, fi *FuncInfo) {
	assert.NotNil(t, fi)
	assert.Equal(t, "Foo", fi.Name)
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

func commonReceiverAsserts(t *testing.T, fi *FuncInfo) {
	assert.NotNil(t, fi.Receivers)
	assert.Len(t, fi.Receivers, 1)
	assert.NotNil(t, fi.Receivers[0])
	assert.Equal(t, "m", fi.Receivers[0].Name)
	assert.Equal(t, "MyStruct", fi.Receivers[0].Ident.Name)
}

func setUp(s string) *ast.FuncDecl {
	i := NewInspector(&InspectorOptions{AvoidNil: true})
	var fn *ast.FuncDecl

	i.FuncDecl.AddListener(func(fd *ast.FuncDecl) {
		fn = fd
	})

	ParseString(i, s)

	return fn
}
