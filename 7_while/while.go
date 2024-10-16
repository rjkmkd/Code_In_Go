package main

import (
	"fmt"
	// "time"
)

func main() {
	// switch statement

	// i := 9
	// switch i {
	// case 1:
	// 	fmt.Println("its one")
	// case 2: 
	// 	fmt.Println("its two")
	// default :
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weekend")

	// default:
	// 	fmt.Println("its weekday")	
	// }


	// type switch

	whoami := func(i interface{}){
		switch i.(type){
		case int: 
			fmt.Println("its an integer")
		case string:
			fmt.Println("its an string")
		
		default :
			fmt.Println("other")

	}
}
whoami(50)
}
