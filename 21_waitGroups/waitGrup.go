package main

import (
	"fmt"
	"sync"
)

// Function `task` takes an `id` and a pointer to `sync.WaitGroup`.
// `task` represents the work each goroutine will do.
// The `defer w.Done()` statement ensures that `Done()` is called once the function completes, decrementing the WaitGroup counter.
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Defer ensures Done is called at the end of this goroutine to signal its completion.
	fmt.Println("doing task", id)
}

func main() {
	// Declare a WaitGroup variable to manage synchronization of goroutines.
	var wg sync.WaitGroup

	// Start 11 goroutines to run `task`.
	for i := 0; i <= 10; i++ {
		wg.Add(1)       // Increment the WaitGroup counter by 1 for each goroutine started
		go task(i, &wg) // Launch a goroutine, passing `i` and a pointer to the WaitGroup
	}

	// Wait until all goroutines have completed by blocking until the counter reaches zero.
	wg.Wait()
}
