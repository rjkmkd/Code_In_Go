package main

import (
	"fmt"
)

// The 'any' keyword can be used as a generic type constraint but it’s generally not recommended
// in cases where we want specific types for better type safety and to avoid unexpected behavior.
// In this code, we are using generics to create functions and types that work with multiple data types,
// applying Go’s type constraints to make them more predictable.

// Uncommented code example using 'any' as a constraint:
/*
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
*/

// A better approach is to specify exactly which types we expect to handle. In this case, 
// we’re specifying that the function can accept slices of int, string, or bool types only.
// Uncommented alternative example specifying specific types manually:
/*
func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
*/

// We are using the 'comparable' constraint in the generic function to ensure that only types
// that support comparison (e.g., int, string) can be used as the generic type `T`.
// Additionally, we define another type parameter `s` for a string, used to pass additional information
// to the function, demonstrating multiple type parameters.
func printSlice[T comparable, s string](items []T, name s) {
	for _, item := range items {
		// Each item from the slice is printed along with the additional name string passed.
		fmt.Println(item, name)
	}
}

// Definition of a generic stack data structure:
// The stack struct is a generic type with a constraint on the allowed types (int, string, bool, float32).
// This ensures we can only use these specific types with our stack struct.
type stack[T int | string | bool | float32] struct {
	elements []T // The 'elements' field holds a slice of elements of type `T`.
}

func main() {
	// Demonstrating usage of printSlice function:
	// Calling printSlice with a slice of booleans and a string "Raj".
	printSlice([]bool{true, false, false, true}, "Raj")

	// Declaring and initializing a stack of type `int`:
	// This stack can only hold integers.
	Mystack := stack[int]{
		elements: []int{1, 3, 4, 4, 5, 57},
	}

	// Declaring and initializing a stack of type `string`:
	// This stack can only hold strings.
	Mystack1 := stack[string]{
		elements: []string{"hdh", "jidjij"},
	}

	// Printing out the contents of both stacks:
	fmt.Println(Mystack)
	fmt.Println(Mystack1)
}
