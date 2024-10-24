
Go Channels

Go channels are a powerful feature that allow goroutines to communicate with each other and synchronize execution. Channels act as conduits through which data flows between goroutines.

Key Concepts:

1. Channel Creation:
   - Channels are created using the `make()` function:
     ch := make(chan T)
     Where `T` is the type of the data the channel will carry. For buffered channels, specify the buffer size:
     ch := make(chan T, bufferSize)

2. Channel Operations:
   - Send: Use the `<-` operator to send data to a channel:
     ch <- value
   - Receive: Use the same operator to receive data from a channel (this is a blocking operation until data is available):
     value := <-ch

3. Buffered vs Unbuffered Channels:
   - Unbuffered channels block the sender until the receiver is ready.
   - Buffered channels allow sending a fixed number of messages without blocking, until the buffer is full.

4. Channel Synchronization:
   - Channels can be used to synchronize tasks. A typical pattern is using a channel to signal when a goroutine has finished a task.
   done := make(chan bool)
   go func() {
       // Perform some work
       done <- true  // Send signal
   }()
   <-done  // Block until signal is received

5. Closing Channels:
   - A channel can be closed to indicate no more values will be sent:
     close(ch)
   - Receivers can detect when a channel is closed by using a second return value in the receive expression:
     value, ok := <-ch
     if !ok {
         // Channel is closed
     }

Code Example Breakdown:

1. Email Sender Example:
   func emailSender(emailChan chan string, done chan bool){
       defer func(){done <- true}()
       for email := range emailChan{
           fmt.Println("sending email to ", email)
           time.Sleep(time.Second)
       }
   }

   func main() {
       emailChan := make(chan string, 100)  // Buffered channel with size 100
       done := make(chan bool)
       
       go emailSender(emailChan, done)  // Start email sender goroutine

       // Send 10 emails
       for i := 0; i < 10; i++ {
           emailChan <- fmt.Sprintf("%d@email.com", i)
       }

       fmt.Println("done sending...")
       close(emailChan)  // Close channel to indicate no more emails
       <-done  // Block until done signal is received
   }

   - Buffered channel: The email channel (emailChan) is buffered with a size of 100, which allows sending multiple emails before the goroutine processes them.
   - Defer: The `defer` statement ensures that `done <- true` is called when the `emailSender` function finishes processing all emails.
   - Channel closing: `close(emailChan)` is used to close the channel once all emails are sent, signaling to the `emailSender` function to stop reading.

2. Task Example:
   func task(done chan bool) {
       defer func() { done <- true }()
       fmt.Println("processing...")
   }

   func main() {
       done := make(chan bool)
       go task(done)  // Start the task in a goroutine
       <-done  // Wait for task completion
   }

   - Channel synchronization: The task function sends a signal via the `done` channel after finishing its work. The main function blocks on `<-done` until the task is complete.

Important Points to Remember:

- Channels are blocking by default. When sending or receiving from an unbuffered channel, the operation will block until the other side is ready.
- Use goroutines to avoid deadlocks when sending and receiving via channels.
- Buffered channels allow some flexibility by letting you send data without blocking until the buffer is full.
- Always close channels when no more data will be sent. However, don’t close a channel from the receiving side or send on a closed channel, as it will cause a panic.
- Defer in goroutines is useful for cleanup tasks or signaling, as it guarantees execution even if the function exits early.
