package main

import (
	"fmt"
	"maps" // The "maps" package provides utility functions for working with maps.
)

func main() {
	// Creating a map using `make`.
	// Syntax: m := make(map[key data type]value data type)
	// m := make(map[string]string) // Declaring a map where both key and value are strings.
	
	// Adding elements to the map using key-value pairs.
	// m["name"] = "goLang"        // Key: "name", Value: "goLang"
	// m["email"] = "ghh@hhv.com"  // Key: "email", Value: "ghh@hhv.com"
	// m["age"] = "30"             // Key: "age", Value: "30"

	// Accessing elements in the map using keys.
	// fmt.Println(m["age"])       // Output: 30 (value associated with key "age")
	// fmt.Println(m["phn"])       // Output: "" (for strings, missing keys return empty string as a zero value)

	// Declaring and initializing a map with values directly.
	// m := map[string]int{"age": 30}
	// fmt.Println(m)              // Output: map[age:30]

	// Checking if a key exists in a map using the `ok` idiom.
	// value, ok := m["age"] // `ok` is true if "age" key exists, else false.
	// if ok {
	//	fmt.Println("ok:", value) // If "age" exists, print the value.
	// } else {
	//	fmt.Println("not ok:", value) // If not, print the zero value of the value type.
	// }

	// Comparing two maps for equality using the `maps.Equal` function.
	m1 := map[string]int{"age": 22, "id": 1}  // Map 1 with "age" and "id" keys.
	m2 := map[string]int{"age": 22, "id": 1}  // Map 2 with the same key-value pairs as Map 1.

	fmt.Println(maps.Equal(m1, m2))           // Output: true (maps are equal if key-value pairs match)
}
