package main

import "fmt"

func main() {
	// Declaring an integer array `nums` of size 4 with default values.
	// var nums [4]int
	// fmt.Println(nums) // Output: [0 0 0 0] (default integer value is 0)
	// nums[0] = 6       // Assigning 6 to the first element of `nums`
	// for i := range nums {
	//	nums[i] = i     // Assign each element to its index value
	// }
	// fmt.Println(nums)  // Output: [0 1 2 3]

	// Declaring a boolean array `vals` of size 4 with default values (false).
	// var vals [4]bool
	// vals[2] = true    // Set the third element to `true`
	// fmt.Println(vals) // Output: [false false true false]

	// Declaring a string array `str` of size 2 with default values (empty strings).
	// var str [2]string
	// fmt.Println(str)  // Output: [  ] (empty strings as default values)
	// str[0] = "hbhh"   // Assign a string to the first element
	// fmt.Println(len(str)) // Output: 2 (length of array is 2)
	// fmt.Println(str)  // Output: [hbhh ]

	// Declaring a float array `vals` of size 2 with default values (0.0).
	// var vals [2]float64
	// fmt.Println(vals) // Output: [0 0] (default float values)

	// Summary of Default Values:
	// - int -> 0
	// - string -> ""
	// - bool -> false
	// - float -> 0.0

	// Declaring and initializing a string array with specific values
	// names := [2]string{"ghg", "gvggh"}
	// fmt.Println(names) // Output: [ghg gvggh]

	// Declaring and initializing a 2D integer array `matrix`
	matrix := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix) // Output: [[1 2] [3 4]]
}
