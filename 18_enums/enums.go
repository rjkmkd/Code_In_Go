package main

import "fmt"

type orderStatus string


const(
	Receved orderStatus = "receved"
	Prepaired = "prepaired"
	Delevered = "delevered"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("changing order status to ",status)
}

// type orderStatus int


// const(
// 	Receved orderStatus = iota
// 	Prepaired
// 	Delevered
// )

// func changeOrderStatus(status orderStatus) {
// 	fmt.Println("changing order status to ",status)
// }


func main() {
	changeOrderStatus(Delevered)
}