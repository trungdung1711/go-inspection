package inspect

import (
	"fmt"
	"go/types"
	"strings"
)

func formatObject(obj types.Object) string {
	if obj == nil {
		return "nil"
	}
	return fmt.Sprintf("%s (%s)", obj.Name(), obj.Type())
}

func PrintScope(scope *types.Scope, depth int) {
	if scope == nil {
		return
	}

	indent := strings.Repeat("    ", depth)
	if depth == 0 {
		fmt.Println("Global Scope (package main)")
	}

	names := scope.Names()
	for _, name := range names {
		obj := scope.Lookup(name)

		// Skip package-level imports like fmt
		if obj.Pkg() != nil && obj.Pkg().Path() == "fmt" {
			continue
		}

		// Print name, type and handle it differently for packages
		fmt.Printf("%s├── %s: %s\n", indent, name, formatObject(obj))

		// Handle struct types
		if typeName, ok := obj.(*types.TypeName); ok {
			underlying := typeName.Type().Underlying()
			if structType, ok := underlying.(*types.Struct); ok {
				fmt.Printf("%s    ├── Struct Fields:\n", indent)
				for j := 0; j < structType.NumFields(); j++ {
					field := structType.Field(j)
					fmt.Printf("%s    │   ├── %s: %s\n", indent, field.Name(), field.Type())
				}
			}
			if iface, ok := underlying.(*types.Interface); ok {
				fmt.Printf("%s    ├── Interface Methods:\n", indent)
				for j := 0; j < iface.NumMethods(); j++ {
					method := iface.Method(j)
					fmt.Printf("%s    │   ├── %s: %s\n", indent, method.Name(), method.Type())
				}
			}
		}

		// Handle functions and methods
		if fn, ok := obj.(*types.Func); ok {
			fmt.Printf("%s    ├── Function Signature: %s\n", indent, fn.Type())
		}

		// Handle nested methods in named types
		if named, ok := obj.Type().(*types.Named); ok {
			for i := 0; i < named.NumMethods(); i++ {
				method := named.Method(i)
				fmt.Printf("%s    ├── Method: %s (%s)\n", indent, method.Name(), method.Type())
			}
		}
	}

	// Handle nested scopes
	for i := 0; i < scope.NumChildren(); i++ {
		PrintScope(scope.Child(i), depth+1)
	}
}
