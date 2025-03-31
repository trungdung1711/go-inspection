package inspect

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func PrintAst(filename string) {
	// Create a new FileSet to manage positions in the source code.
	fs := token.NewFileSet()

	// Parse the source code (src) into an abstract syntax tree (AST).
	// The parser.ParseFile function takes the FileSet, a filename (empty here),
	// the source code, and a mode (parser.AllErrors to capture all errors).
	node, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		// Print an error message if parsing fails and exit the function.
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the AST structure to the standard output.
	// ast.Print traverses the AST and displays its nodes in a readable format.
	ast.Print(fs, node)
}
