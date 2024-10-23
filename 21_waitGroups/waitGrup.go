package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()  // When this goroutine completes, it calls Done() to decrement the counter
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup  // Declare a WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)          // Increment the WaitGroup counter by 1 for each goroutine
		go task(i, &wg)    // Start a goroutine for each task, passing the WaitGroup as a pointer
	}

	wg.Wait()  // Block the main goroutine until the counter becomes 0 (all tasks are done)
}
