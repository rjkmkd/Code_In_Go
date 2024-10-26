package main

import (
	"fmt"
	"time"
)

// Declaring an anonymous function within a goroutine for concurrent execution.
// Uncommented example function for performing a task:
/*
func task(id int) {
	id *= 2
	fmt.Println("doing task", id)
}
*/

func main() {
	// Launching goroutines in a loop:
	// Here, a loop spawns 11 goroutines, each printing the value of `i` (from 0 to 10).
	// Goroutines are lightweight threads managed by the Go runtime, allowing concurrent execution.
	for i := 0; i <= 10; i++ {
		go func(i int) {
			// Anonymous function that takes `i` as an argument.
			// Each goroutine prints the current value of `i` passed from the loop.
			fmt.Println(i)
		}(i)
	}

	// Adding a sleep to prevent main from exiting too early:
	// Without `time.Sleep`, the main function may exit before the goroutines have a chance to complete their tasks.
	time.Sleep(time.Second * 2)
}
