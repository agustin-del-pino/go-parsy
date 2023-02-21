package main

import (
	"fmt"
	"go-parsy/pkg/parsy"
	"go/ast"
)

func main() {
	ins := parsy.NewInspector(&parsy.InspectorOptions{
		AvoidNil: true,
	})

	ins.All.AddListener(func(n ast.Node) {
		fmt.Printf("%T\n", n)
	})

	parsy.ParseString(ins, `package main
	var (
		num = 3
		str = "Hello World"
	)
	`)
}
