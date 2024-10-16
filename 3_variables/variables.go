package main

import "fmt"

func main() {
	// variables
	// we have to use this syntax when we declare varible first and then after sometime as per requirement we assign the value on it.
	// var name string  = "golang"
	var name string ;
	name = "golang"
	fmt.Println(name)
	// short note
	// if you are not declared the datatype then but assign the value of the variable at the time of declaration, it automatically detect the datatype .
	var age = 30    //int
	fmt.Println(age)  
	var isAdult = true //bool
	fmt.Println(isAdult)

	// short hand syntax
	// only use this when we declare a variable and assign value at the same time
	email := "raj@gmail.com"
	fmt.Println(email)

	//summary
	var price float32 = 50.5
	// var price = 50.5
	// price:=50.8
	fmt.Println(price)
}