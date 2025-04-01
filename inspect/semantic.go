package inspect

import (
	"fmt"
	"go/types"
	"strings"
)

// printObject prints the basic and important details of a single object in the scope in one line.
func printObject(indent string, name string, obj types.Object) {
	// Prepare the type of the object
	objType := obj.Type()

	// Prepare additional details about the object
	var typeDetails string

	// Switch to determine the type of the object and provide specific details
	switch t := objType.Underlying().(type) {
	case *types.Signature:
		// Function signature
		typeDetails = fmt.Sprintf("function, returns: %s", t.Results())
	case *types.Struct:
		// Struct type, print the number of fields
		typeDetails = fmt.Sprintf("struct with %d field(s)", t.NumFields())
	case *types.Basic:
		// Basic types (int, string, etc.)
		typeDetails = fmt.Sprintf("basic type, %s", t.String())
	case *types.Array:
		// Array type
		typeDetails = fmt.Sprintf("array of %s with length %d", t.Elem(), t.Len())
	case *types.Slice:
		// Slice type
		typeDetails = fmt.Sprintf("slice of %s", t.Elem())
	case *types.Map:
		// Map type
		typeDetails = fmt.Sprintf("map with key type %s and value type %s", t.Key(), t.Elem())
	case *types.Named:
		// Named types (e.g., struct, interface)
		typeDetails = fmt.Sprintf("named type, %s", t.Obj().Name())
	default:
		// Default case for other types
		typeDetails = objType.String()
	}

	// Print object details in one line
	fmt.Printf("%s├── %s: %s (%s)\n", indent, name, objType, typeDetails)
}

// PrintScope recursively prints the scope tree and its contained elements in a tree-like structure with more detailed information.
func PrintScope(scope *types.Scope, depth int) {
	if scope == nil {
		return
	}

	// Indentation to show the depth in the scope hierarchy
	indent := strings.Repeat("│   ", depth)

	// Print the current scope's name
	fmt.Printf("%sScope (Level %d):\n", indent, depth)

	// Store the names of objects within the scope
	names := scope.Names()

	// Collect the objects and print them with more detailed information
	for i, name := range names {
		obj := scope.Lookup(name)

		// Check if it's the last element in the scope level
		if i == len(names)-1 {
			// Last item in scope, no separator at the end
			printObject(indent, name, obj)
		} else {
			// Not the last item, use a separator
			printObject(indent, name, obj)
		}
	}

	// Recursively handle nested scopes (child scopes)
	for i := 0; i < scope.NumChildren(); i++ {
		childScope := scope.Child(i)
		if childScope != nil {
			// Print child scopes recursively
			PrintScope(childScope, depth+1)
		}
	}
}
