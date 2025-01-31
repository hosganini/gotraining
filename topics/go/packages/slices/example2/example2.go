// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program shows how to use the Clip API from the slices package.
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Clip removes the unused capacity from the slice.

func main() {
	list := make([]string, 0, 10)
	fmt.Printf("Len(%d), Cap(%d)\n", len(list), cap(list))

	// -------------------------------------------------------------------------
	// Append a string to the slice

	list = append(list, "A")
	fmt.Printf("Len(%d), Cap(%d)\n", len(list), cap(list))

	// -------------------------------------------------------------------------
	// Clip

	list = slices.Clip(list)
	fmt.Printf("Len(%d), Cap(%d)\n", len(list), cap(list))
}
