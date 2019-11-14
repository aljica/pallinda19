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
		ch <- i
	}

	close(ch)
	waitgroup.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, waitgroup *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
	waitgroup.Done()
}
