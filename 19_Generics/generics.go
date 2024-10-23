package main

import (
	"fmt"
)
	// but any is note a good practice to use.
// func printSlice[T any](items [] T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

	// we can put manualy the dadatypes which we may pass
// func printSlice[T int | string | bool](items [] T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// we use comparable interface
// also we can pass multiple arguments
func printSlice[T comparable, s string ](items [] T, name s) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack [T int|string|bool|float32] struct{
	elements [] T
}
func main() {
	// printSlice([]int{4,5,6})
	// printSlice([]string{"hduh","hdhu","uhhduh"})
	printSlice([]bool{true, false, false, true},"Raj")


	Mystack := stack[int]{
		elements: []int{1,3,4,4,5,57},
	}
	Mystack1 := stack[string]{
		elements: []string{"hdh","jidjij"},
	}
	fmt.Println(Mystack)
	fmt.Println(Mystack1)
}