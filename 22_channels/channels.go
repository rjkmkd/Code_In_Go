package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

// func processNum(numChan chan int){
// 	for num:= range numChan{
// 		fmt.Println("processing number: ",num)
// 		time.Sleep(time.Second)
// 	}

// }

// func sun(result chan int, num1 int, num2 int){
// 	sumResult := num1 + num2
// 	result <-sumResult

// }

// func task(done chan bool){
// 	defer func(){done <- true}()
// 	fmt.Println("processing...")

// }

func emailSender(emailChan chan string, done chan bool){
	defer func (){done<-true}()
	for email:=range emailChan{
		fmt.Println("sending email to ",email)
		time.Sleep(time.Second)
	}
}

func main() {

	emailChan :=make(chan string, 100)
	done:=make(chan bool)
	go emailSender(emailChan,done)
	for i:=0;i<10;i++{
		emailChan<-fmt.Sprintf("%d@email.com",i)
	}

	fmt.Println("done sending...")
	close(emailChan)
	<-done
	// emailChan <-"1@example.com"
	// emailChan <-"2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)


	// done :=make(chan bool)
	// go task(done)
	// <-done //block



	// result := make(chan int)

	// go sun(result, 23,3)
	// res:=<-result  //blocking

	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)

	// messageChan := make(chan string)
	// messageChan <- "ping" //blocking
	// msg := <-messageChan
	// fmt.Println(msg)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	
}