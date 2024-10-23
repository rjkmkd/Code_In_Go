package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	id*=2
// 	fmt.Println("doing task",id)
// }
func main() {
	// for i:=0;i<=10;i++{
	// 	go task(i)
	// }
	for i:=0;i<=10;i++{
		go func(i int){
		fmt.Println(i)
	}(i)
	}

	time.Sleep(time.Second*2)
}