package parsy

import (
	"go/parser"
	"go/token"
)

func parse(i *Inspector, p string, s any) error {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, p, s, parser.AllErrors)

	if err != nil {
		return err
	}

	InspectNode(i, f)

	return nil
}

func ParseString(i *Inspector, s string) error {
	return parse(i, "", s)
}

func ParseFile(i *Inspector, p string) error {
	return parse(i, p, nil)
}
