package main

import "fmt"

// Declaring an interface `paymenter`:
// Interfaces in Go act as contracts. Any type implementing this interface must provide a `pay` method with the specified signature (taking a `float32` amount argument).
type paymenter interface {
	pay(amount float32) // Method that each type implementing `paymenter` must define
	/*
	Explanation:
	This interface specifies a contract for any struct that implements it. Any struct implementing the `paymenter` interface
	must have a `pay` method that accepts a `float32` argument `amount`. This ensures that we can use different payment methods
	with the same interface, enabling flexibility and polymorphism.
	*/
}

// Struct `payment`:
// This struct has a field `gateway` of type `paymenter`, allowing `payment` to use any struct that implements the `paymenter` interface.
type payment struct {
	// gateway could be any payment gateway that implements the `paymenter` interface.
	gatway paymenter
}

// Struct `razorPay` implementing the `paymenter` interface:
type razorPay struct{}

// Method `pay` for `razorPay` struct:
// Defines the `pay` method required by the `paymenter` interface. This is specific to RazorPay and contains logic for payment using RazorPay.
func (r razorPay) pay(amount float32) {
	// Logic to make payment using RazorPay
	fmt.Println("making payment using RazorPay", amount)
}

// Method `makePayment` for the `payment` struct:
// The `makePayment` method accepts an `amount` and calls the `pay` method of the assigned `gateway`. 
// This makes it flexible to switch between different payment gateways at runtime.
func (p payment) makePayment(amount float32) {
	p.gatway.pay(amount)
}

// Struct `strype` implementing the `paymenter` interface:
type strype struct{}

// Method `pay` for `strype` struct:
// Implements the `pay` method required by the `paymenter` interface for Stripe payments.
func (s strype) pay(amount float32) {
	fmt.Println("making payment using Stripe", amount)
}

// Struct `fakepaymentGW` implementing the `paymenter` interface:
// Useful for testing, as it simulates a payment gateway without actual transaction logic.
type fakepaymentGW struct{}

func (f fakepaymentGW) pay(amount float32) {
	fmt.Println("making payment using fake payment gateway", amount)
}

// Struct `paypal` implementing the `paymenter` interface:
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("payment with PayPal:", amount)
}

func main() {
	// Example: Initializing `payment` with the `paypal` payment gateway.
	// Any payment gateway that implements the `paymenter` interface can be assigned here.
	myPaypalPaymentGw := paypal{}
	MyPayment := payment{
		gatway: myPaypalPaymentGw,
	}

	// Calling `makePayment` method, which will use the specified `gatway` to process the payment.
	MyPayment.makePayment(100)
}
