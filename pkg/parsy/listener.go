package parsy

import "go/ast"

// Listener type is the callback that it 
// will execute when the T node is detected
type Listener[T ast.Node] func(T)

type listeners[T ast.Node] struct {
	lsts []Listener[T]
}

// AddListener adds a new listener
func (ls *listeners[T]) AddListener(l Listener[T]) {
	ls.lsts = append(ls.lsts, l)
}

func (ls *listeners[T]) trigger(n T) {
	for _, l := range ls.lsts {
		l(n)
	}
}