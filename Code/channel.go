package main

import "fmt"

func channel() {

	messages := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	go func() { messages <- "ping" }() //Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	msg := <-messages //The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	fmt.Println(msg)
}

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

// Go channels are used for communicating between concurrently running functions by sending and receiving a specific element type's data. When we have numerous Goroutines running at the same time, channels are the most convenient way for them to communicate with one another.

// Goroutines provide a means for multiprocessing in a Golang application, allowing multiple processes to run simultaneously. In addition, Channels and WaitGroups enable passing data between threads or blocking one thread until another completes.
