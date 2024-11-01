package main

import (
	// "fmt"
	// "slices" // Importing "slices" package for slice utility functions like Equal
)

func main() {
	// A slice with no initial values is `nil` and has length 0.
	// var nums []int
	// fmt.Println(nums)       // Output: [] (nil slice)
	// fmt.Println(len(nums))   // Output: 0 (length of the slice)

	// Another way to initialize a slice with `make` (initial length: 0, capacity: 3).
	// var nums = make([]int, 0, 3)

	// Initializing and using a slice with values.
	// nums := []int{1, 2, 3, 4}      // `nums` is a slice with four elements.
	// nums = append(nums, 6)         // Appending an element increases the slice's length.
	// nums[4] = 3                    // Directly modifying an element in the slice
	// fmt.Println(nums)              // Output: [1 2 3 4 3]
	// fmt.Println(len(nums))         // Output: 5 (current length of the slice)
	// fmt.Println(cap(nums))         // Output: capacity of the slice, dynamically managed by Go

	// Using the slice operator `[:]` to extract portions of a slice.
	// var nums []int
	// nums = []int{1, 2, 3, 4}
	// fmt.Println(nums[0:])          // Output: [1 2 3 4] (entire slice from index 0 onward)
	// fmt.Println(nums[0:2])         // Output: [1 2] (slice from index 0 up to, but not including, index 2)

	// Using the `slices` package (commented out because "slices" is not standard).
	// The `slices` package provides functions like `Equal` to compare slices.
	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 3}
	// fmt.Println(slices.Equal(num1, num2)) // Output: true (both slices are equal)

	// Copying elements from one slice to another.
	// var num1 = []int{1, 2, 3}       // Source slice
	// var num2 = make([]int, len(num1)) // Destination slice with the same length as `num1`
	// copy(num2, num1)                // Copying elements from `num1` to `num2`
	// fmt.Println(num1)               // Output: [1 2 3]
	// fmt.Println(num2)               // Output: [1 2 3] (num2 has the same elements as num1)
}
