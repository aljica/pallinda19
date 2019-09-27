package main

import (
	"fmt"
	"sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	ch := make(chan int)

	go func() {
		Print(ch)
		waitgroup.Done()
	}()

	go func() {
		for i := 1; i <= 11; i++ {
			ch <- i
		}
		close(ch)
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
