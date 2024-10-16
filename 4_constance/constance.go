package main

import "fmt"

//note:-
// age :=30 this syntax is invalid out side function.
// var age = 30 we can use this syntax
// var age int = 30 we can use this syntax also
// const name = "golang" also use const outside the function
func main(){
	const name string = "golang"
	// can't reassign the value
	// name = "java"

	// constant grouping :- when we have to declare multiple constants, so we can group them.
	const (
		port = 5000
		host = "localhost"
	)
	// post = 5500 // give me error
	var(
		name67 = "raj"
		age67 = "22"
	)
	name67 = "Bono"
	fmt.Println(name67, age67)
	fmt.Println(port, host)
}