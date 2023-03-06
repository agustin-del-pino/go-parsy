package parsy

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAs_with_valid_casting(t *testing.T) {
	var n ast.Node

	n = &ast.BasicLit{}

	l, ok := As[*ast.BasicLit](n)

	assert.IsType(t, &ast.BasicLit{}, l)
	assert.True(t, ok)

}

func TestAs_with_invalid_casting(t *testing.T) {
	var n ast.Node

	n = &ast.BasicLit{}

	l, ok := As[*ast.Ident](n)

	assert.Nil(t, l)
	assert.False(t, ok)

}
