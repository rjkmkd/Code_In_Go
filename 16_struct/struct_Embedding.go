package main

import "fmt"

type consumer struct {
	name   string
	con_id int
}

type Order struct {
	id     int
	amount float32
	status string
	consumer
}

func main() {
	// newConsumer := consumer{
	// 	name:   "Raj",
	// 	con_id: 678,
	// }

	// myOrder := Order{
	// 	id:       123,
	// 	amount:   39.99,
	// 	status:   "pending",
	// 	consumer: newConsumer,
	// }

	//another syntax

	myOrder := Order{
		id:       123,
		amount:   39.99,
		status:   "pending",
		consumer:consumer{
			name: "Raj",
			con_id: 6789,
		},
	}
	fmt.Println(myOrder.consumer)

}