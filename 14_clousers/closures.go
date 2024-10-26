package main

import "fmt"

// Main function to execute the program logic.
func main() {
	// `getCounter` is assigned the function returned by `counter`.
	// This returned function will retain access to the `count` variable inside `counter`.
	getCounter := counter()

	// Each call to `getCounter()` increments the `count` variable by 1 and returns its value.
	fmt.Println(getCounter()) // Output: 1
	fmt.Println(getCounter()) // Output: 2
	fmt.Println(getCounter()) // Output: 3
	fmt.Println(getCounter()) // Output: 4
}

// `counter` function returns an anonymous function that increments and returns an internal `count` variable.
func counter() func() int {
	// Initialize `count` with zero. This variable is scoped to the `counter` function.
	count := 0

	// Return a function that increments and returns `count`.
	// This anonymous function has a closure over `count`, meaning it can access and modify `count` even after `counter` has returned.
	return func() int {
		count += 1 // Increment `count` by 1 each time this function is called.
		return count // Return the updated `count`.
	}
}
