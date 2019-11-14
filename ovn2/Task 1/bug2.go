package main

import (
	"fmt"
	"sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	ch := make(chan int)

	go Print(ch, &waitgroup)

	for i := 1; i <= 11; i++ {
		ch <- i // This line waits until Print() is ready to receive
		// Once Print() receives, the for-loop moves on to the next iteration.
	}

	close(ch) // Because of the above (see comments), we can instantly
	// close the channel upon exiting the for-loop, because it means
	// we have transmitted all of our data through the channel ch.
	waitgroup.Wait() // Here we just wait for Print() to finish its final Println.
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, waitgroup *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
	waitgroup.Done()
}
