package main

import (
	"fmt"
	// "structs"
	// "time"
)

// Order struct
//  type order struct{
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time
//  }

//  //constructor

//  func NewOrder (id string , amount float32, status string) *order{
// 	myOrder := order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 		createdAt: time.Now(),
// 	}
// 	return &myOrder
//  }

//  func (o *order) updateStatus(newStatus string){
// 	o.status = newStatus
//  }
// // if you want to cahnge value inside the struct instance , then you have to use pointer.
//  func (o *order) changeStatus(status string){
// 	o.status = status
//  }
// // if you want to get any value from the struct instance/ return value then you may not use pointer.
//  func (o order) getAmmount()float32{
// 	return o.amount
//  }

// func main(){

	// myOrder := order{
	// 	id:"1",
	// 	amount: 300.99,
	// 	status: "received",
	// 	createdAt: time.Now(),
	// }
	// myOrder2 := order{
	// 	id:"2",
	// 	amount: 500.99,
	// 	status: "received",
	// 	createdAt: time.Now(),
	// }
	// myOrder2.status = "shipped"
	// myOrder.changeStatus("confirmed")
	// fmt.Println("myOrder: ",myOrder)
	// fmt.Println("myOrder2: ",myOrder2)
	// fmt.Println(myOrder.status)
	// fmt.Println("my order ammount is : ",myOrder.getAmmount())


	// using constructor

	// myOrder1 :=NewOrder("1",30.98,"received")
	// fmt.Println(*myOrder1)
	// fmt.Println(myOrder1.status)
	// myOrder1.updateStatus("pending")
	// fmt.Println(*myOrder1)
	// fmt.Println(myOrder1.status)

	//struct using short hand notation

	language:=struct{
		name string
		isGood bool
	}{"golang",true}
	fmt.Println("language: ",language)

// }