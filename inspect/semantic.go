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

	// Prepare additional details about the object
	var typeDetails string

	// Switch to determine the type of the object and provide specific details
	switch t := objType.Underlying().(type) {
	case *types.Signature:
		// Function signature: print parameters and return types
		params := t.Params()
		results := t.Results()
		paramStr := []string{}
		for i := 0; i < params.Len(); i++ {
			paramStr = append(paramStr, params.At(i).Type().String())
		}
		resultStr := []string{}
		for i := 0; i < results.Len(); i++ {
			resultStr = append(resultStr, results.At(i).Type().String())
		}
		// Print function details
		fmt.Printf("%s├── %s: func()\n", indent, name)
		fmt.Printf("%s│   ├── Parameters: [%s]\n", indent, strings.Join(paramStr, ", "))
		fmt.Printf("%s│   └── Returns: [%s]\n", indent, strings.Join(resultStr, ", "))
		return
	case *types.Struct:
		// Struct type: print the number of fields and their types
		fieldDetails := []string{}
		for i := 0; i < t.NumFields(); i++ {
			field := t.Field(i)
			fieldDetails = append(fieldDetails, fmt.Sprintf("%s: %s", field.Name(), field.Type()))
		}
		// Print struct details
		fmt.Printf("%s├── %s: struct\n", indent, name)
		for _, fieldDetail := range fieldDetails {
			fmt.Printf("%s│   └── %s\n", indent, fieldDetail)
		}
		return
	case *types.Basic:
		// Basic types (int, string, etc.)
		typeDetails = fmt.Sprintf("basic type, %s", t.String())
		// Print basic type details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	case *types.Array:
		// Array type: print length and element type
		typeDetails = fmt.Sprintf("array of %s with length %d", t.Elem(), t.Len())
		// Print array details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	case *types.Slice:
		// Slice type: print element type
		typeDetails = fmt.Sprintf("slice of %s", t.Elem())
		// Print slice details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	case *types.Map:
		// Map type: print key and value types
		typeDetails = fmt.Sprintf("map with key type %s and value type %s", t.Key(), t.Elem())
		// Print map details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	case *types.Named:
		// Named types (e.g., struct, interface)
		typeDetails = fmt.Sprintf("named type, %s", t.Obj().Name())
		// Print named type details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	case *types.Interface:
		// Interface type: print the methods it defines
		methodDetails := []string{}
		for i := 0; i < t.NumMethods(); i++ {
			method := t.Method(i)
			methodDetails = append(methodDetails, fmt.Sprintf("%s: %s", method.Name(), method.Type()))
		}
		// Print interface details with methods
		fmt.Printf("%s├── %s: interface\n", indent, name)
		for _, methodDetail := range methodDetails {
			fmt.Printf("%s│   └── %s\n", indent, methodDetail)
		}
		return
	default:
		// Default case for other types
		typeDetails = objType.String()
		// Print default type details
		fmt.Printf("%s├── %s: %s\n", indent, name, typeDetails)
		return
	}
}

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
