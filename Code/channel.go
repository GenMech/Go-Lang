package main

import "fmt"

func channel() {

	messages := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	go func() { messages <- "ping" }() //Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	msg := <-messages //The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	fmt.Println(msg)
}

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
