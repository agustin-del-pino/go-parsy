package parsy

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListeners_AddListener_with_no_array_init(t *testing.T) {
	ls := &listeners[ast.Node]{}

	ls.AddListener(func(n ast.Node) {})

	assert.Len(t, ls.lsts, 1)
}

func TestListeners_trigger_with_no_listener(t *testing.T) {
	out := false
	ls := &listeners[ast.Node]{}
	n := ast.NewIdent("")

	ls.trigger(n)

	assert.False(t, out)
}

func TestListeners_trigger_with_a_listener(t *testing.T) {
	out := false
	ls := &listeners[ast.Node]{
		lsts: []Listener[ast.Node]{func(n ast.Node) {
			out = true
		}},
	}
	n := ast.NewIdent("")

	ls.trigger(n)

	assert.True(t, out)
}

func TestListeners_trigger_with_more_than_one_listener(t *testing.T) {
	out := 0
	ls := &listeners[ast.Node]{
		lsts: []Listener[ast.Node]{
			func(n ast.Node) {out++}, func(n ast.Node) {out++}, func(n ast.Node) {out++}},
	}
	n := ast.NewIdent("")

	ls.trigger(n)

	assert.Equal(t, 3, out)
}
