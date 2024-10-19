package main

import "fmt"

//interface => basically a contarct
type paymenter interface{
	pay(amount float32)// method 

	/*explaination => this means which ever struct implement this interface, mandatory to have a method name "pay" which takes argument "amount" of "float32" type.
	*/

}

type payment struct{
	// gatway strype
	// gatway razorPay
	gatway paymenter
}

type razorPay struct{}

func (p payment) makePayment(amount float32) {
	// razorPayPaymentGW:=razorPay{}
	// razorPayPaymentGW.pay(amount)

	// StripePaymentGW:=strype{}
	// StripePaymentGW.pay(amount)
	p.gatway.pay(amount)
}

func (r razorPay) pay(amount float32) {
	//logic to make payment

	fmt.Println("making payment using razorPay ",amount )
}

type strype struct{}

func (s strype) pay(amount float32){
	fmt.Println("making payment using stripe ",amount )
}

//for testing

type fakepaymentGW struct{}
func (f fakepaymentGW) pay(amount float32){
	fmt.Println("making payment using fake payment gateway ",amount )
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("payment with paypal: ",amount)
}

func main() {
	// stripePaymentGW:=strype{}
	// MyPayment := payment{
	// 	gatway:stripePaymentGW ,
	// }

	// razorPayPaymentGW:=razorPay{}
	// MyPayment :=payment{
	// 	gatway: razorPayPaymentGW,
	// }
	// MyfakepatmentGW :=fakepaymentGW{}
	// MyPayment:= payment{
	// 	gatway:MyfakepatmentGW,
	// }

	// paypal
	myPaypalPaymentGw:=paypal{}
	MyPayment:=payment{
		gatway: myPaypalPaymentGw,
	}

	MyPayment.makePayment(100)
}