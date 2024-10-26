Explanation of Concepts in the Code
Goroutines:

Goroutines are created using the go keyword before a function call. This allows the function to run concurrently.
Each iteration of the loop spawns a new goroutine, executing an anonymous function that prints the current value of i.
Anonymous Function with Parameters:

The function func(i int) is an anonymous (or inline) function that accepts an integer i.
Passing i as an argument to the anonymous function (i.e., (i)) captures its value for each goroutine, ensuring that each goroutine prints a consistent value for i.
Concurrency Management:

Using time.Sleep(time.Second * 2) gives the goroutines enough time to complete before the main function exits.
Without the sleep, the program would likely exit before any goroutines complete their tasks, as Go does not wait for goroutines to finish by default.
Race Conditions:

The sleep method is a basic approach and not ideal for managing goroutine completion in real applications.
For better concurrency control, synchronization techniques like sync.WaitGroup can be used to manage goroutine execution.