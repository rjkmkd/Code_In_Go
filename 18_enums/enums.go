package main

import "fmt"

// Defining a custom type `orderStatus` as a string.
// This allows us to define constants of this custom type to represent specific order statuses.
type orderStatus string

// Defining constants for the possible order statuses:
// Here, `Receved`, `Prepaired`, and `Delevered` are constants of the type `orderStatus`.
// The values are strings, which means we can print these statuses as human-readable strings in the output.
const (
	Receved   orderStatus = "receved"   // Status representing "Received"
	Prepaired orderStatus = "prepaired" // Status representing "Prepared"
	Delevered orderStatus = "delevered" // Status representing "Delivered"
)

// Function `changeOrderStatus` accepts a parameter of type `orderStatus` and prints the status.
// This function demonstrates how to work with custom types and ensures that only values of type `orderStatus` are passed.
func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to", status)
}

// Below is an alternative approach to define `orderStatus` as an integer type and use `iota` for consecutive values:
// By using integers for `orderStatus`, each status is represented by a unique integer value.
// Uncommented alternative integer-based approach:
/*
type orderStatus int

const (
	Receved orderStatus = iota // Starts with 0
	Prepaired                  // Value of 1
	Delevered                  // Value of 2
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to", status)
}
*/

func main() {
	// Calling the `changeOrderStatus` function with `Delevered` as the status.
	// This will output: "changing order status to delevered".
	changeOrderStatus(Delevered)
}
