# go-parsy
Go Parser wrapper for more easy and understandable approach.

# Overview

````go
ins := parsy.NewInspector(&parsy.InspectorConfig{
  AvoidNil: true
})

ins.All.AddListener(func (n ast.Node){
  fmt.Printf("%T\n", n)
})

parsy.ParseString(ins, "package main; var num = 3")

// Output:
// *ast.File
// *ast.Ident
// *ast.GenDecl
// *ast.ValueSpec
// *ast.Ident
// *ast.BasicLit
````

# Inspector
The **Inspector** is the listeners container of any kind of node, even the *all* nodes *(any node no matter of its kind)*.

In order to add a new listener to a *Node Type*, select the wanted type follow by `AddListener` method. 

The `AddListener` method requires a **Listener Callback** defined as:
````go
type Listener[T ast.Node] func(T)
````
Where `T` it will the *Node Type*, because of that the param will be the actual node casted to the specific *Node Type*.

## Inspector Options
- `AvoidNil`: ignore the Nodes that are `nil`.

# Parser Function
These functions are wrapper of a internal parser function that look like this:

````go
func parse(i *Inspector, p string, s any) error {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, p, s, parser.AllErrors)

	if err != nil {
		return err
	}

	ast.Walk(i, f)

	return nil
}
````

## ParseString
This function will parse a source code from string and executes the inspector.

**Requires**: an inspector pointer and the actual source code as string.

**Returns**: an `error` in case of unsuccessful parse, otherwise `nil`.

````go
parsy.ParseString(ins, "package main; var num = 3")
````

## ParseFile
This function will parse a source code from file and executes the inspector.

**Requires**: an inspector pointer and the actual source code from a filepath.

**Returns**: an `error` in case of unsuccessfully parse, otherwise `nil`.
````go
parsy.ParseFile(ins, "./main.go")
````

# InspectNode
This function will inspect any `ast.Node` by executing the given inspector.

**Requires**: an inspector pointer and the actual node.

````go
parsy.ParseNode(ins, n)
````

