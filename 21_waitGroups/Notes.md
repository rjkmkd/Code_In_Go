Explanation of Concepts in the Code
WaitGroup (sync.WaitGroup):

sync.WaitGroup is used to synchronize the completion of multiple goroutines.
It allows the main function to wait for all goroutines to complete before it exits.
Each Add(1) call increments the WaitGroup counter, and each Done() call decrements it. The program blocks at Wait() until the counter reaches zero.
Using Add, Done, and Wait with WaitGroup:

wg.Add(1): Called before starting each goroutine, this increments the WaitGroup counter to track the task.
defer w.Done(): This statement ensures that Done() is called at the end of each goroutine, signaling task completion.
wg.Wait(): This blocks the main function until all goroutines have called Done().
Concurrency with Goroutines:

Each iteration of the loop starts a new goroutine using the task function, where i is passed as an identifier.
Using WaitGroup, the program manages these concurrent tasks without prematurely exiting the main function.
Ensuring Proper Synchronization:

Without sync.WaitGroup, there would be a race condition where main might exit before all goroutines complete.
This WaitGroup-based synchronization is efficient and avoids the need for time.Sleep, which could result in inconsistent results in concurrent programs.