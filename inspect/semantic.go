package inspect

import (
	"fmt"
	"go/types"
	"strings"
)

// printObject prints the details of a single object in the scope in a tree-like structure.
func printObject(indent string, name string, obj types.Object) {
	// Prepare the type of the object
	objType := obj.Type()
	underlyingType := objType.Underlying()

	// Switch to determine the type of the object and provide specific details
	switch t := objType.(type) {
	case *types.Signature:
		// Function signature: print parameters and return types
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   ├── Underlying: %T\n", indent, underlyingType)

		// Print parameters
		if t.Params().Len() > 0 {
			fmt.Printf("%s│   ├── Parameters:\n", indent)
			for i := 0; i < t.Params().Len(); i++ {
				param := t.Params().At(i)
				fmt.Printf("%s│   │   ├── %s: %s\n", indent, param.Name(), param.Type())
			}
		}

		// Print return values
		if t.Results().Len() > 0 {
			fmt.Printf("%s│   └── Returns:\n", indent)
			for i := 0; i < t.Results().Len(); i++ {
				result := t.Results().At(i)
				fmt.Printf("%s│       ├── %s\n", indent, result.Type())
			}
		}
		return
	case *types.Struct:
		// Struct type: print the fields
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   ├── Underlying: %T\n", indent, underlyingType)

		if t.NumFields() > 0 {
			fmt.Printf("%s│   └── Fields:\n", indent)
			for i := 0; i < t.NumFields(); i++ {
				field := t.Field(i)
				fmt.Printf("%s│       ├── %s: %s\n", indent, field.Name(), field.Type())
			}
		}
		return
	case *types.Basic:
		// Print basic type details
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   └── Underlying: %T\n", indent, underlyingType)
		return
	case *types.Array:
		// Print array type details
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   └── Underlying: %T\n", indent, underlyingType)
		return
	case *types.Slice:
		// Print slice type details
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   └── Underlying: %T\n", indent, underlyingType)
		return
	case *types.Map:
		// Print map type details
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   └── Underlying: %T\n", indent, underlyingType)
		return
	case *types.Named:
		// Print named type details
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   ├── Underlying: %T\n", indent, underlyingType)

		// If the underlying type is a struct, print its fields
		if structType, ok := underlyingType.(*types.Struct); ok {
			fmt.Printf("%s│   ├── Fields:\n", indent)
			for i := 0; i < structType.NumFields(); i++ {
				field := structType.Field(i)
				fmt.Printf("%s│   │   ├── %s: %s\n", indent, field.Name(), field.Type())
			}
		}

		// Print methods of the named type
		if t.NumMethods() > 0 {
			fmt.Printf("%s│   └── Methods:\n", indent)
			for i := 0; i < t.NumMethods(); i++ {
				method := t.Method(i)
				fmt.Printf("%s│       ├── %s: %s\n", indent, method.Name(), method.Type())
			}
		}
		return
	case *types.Interface:
		// Interface type: print the methods it defines
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   ├── Underlying: %T\n", indent, underlyingType)

		if t.NumMethods() > 0 {
			fmt.Printf("%s│   └── Methods:\n", indent)
			for i := 0; i < t.NumMethods(); i++ {
				method := t.Method(i)
				fmt.Printf("%s│       ├── %s: %s\n", indent, method.Name(), method.Type())
			}
		}
		return
	default:
		// Default case for other types
		fmt.Printf("%s├── %s [Object Type: %T]\n", indent, name, obj)
		fmt.Printf("%s│   ├── Type: %T\n", indent, objType)
		fmt.Printf("%s│   └── Underlying: %T\n", indent, underlyingType)
		return
	}
}

// print the scope in a tree-like structure
func PrintScope(scope *types.Scope, depth int) {
	if scope == nil {
		return
	}

	// Indentation to show the depth in the scope hierarchy
	indent := strings.Repeat("    ", depth)

	// Print the current scope's level
	if depth > 0 {
		fmt.Printf("%s├── Scope (Level %d):\n", indent, depth)
	} else {
		fmt.Printf("%sScope (Level %d):\n", indent, depth)
	}

	// Print the objects in the current scope
	names := scope.Names()
	for _, name := range names {
		obj := scope.Lookup(name)
		// Print each object; adjust the connector as needed
		printObject(indent+"│   ", name, obj)
	}

	// Recursively print child scopes
	for i := 0; i < scope.NumChildren(); i++ {
		// childIndent := indent + "│   " // Consistent indentation for children
		PrintScope(scope.Child(i), depth+1)
	}
}
