package inspect

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// PrintAst parses the Go source code from the specified file and prints its
// abstract syntax tree (AST) structure to the standard output. This function
// is useful for inspecting the internal representation of Go code.
//
// Parameters:
//   - filename: The path to the Go source file to be parsed.
//
// If the file cannot be parsed, the function prints an error message and
// terminates the program.
func PrintAst(filename string) {
	fs := token.NewFileSet()

	node, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ast.Print(fs, node)
}
