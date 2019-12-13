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

	close(ch) /* Because of the above (see comments), we can instantly
	close the channel upon exiting the for-loop, because it means
	we have transmitted all of our data through the channel ch.*/
	waitgroup.Wait() /* Here we just wait for Print() to finish its final Println.
	Prior to adding the waitgroup, we had a data race. As soon as i = 11 was sent 
	through the channel, we instantly closed the channel. This meant that, while 
	"for n:=range ch" in Print() received i=11, we never actually waited for 
	the program to get to the Println(). So, sometimes, the Println() was faster 
	than the close(ch), and other times it wasn't, essentially creating a data race 
	in which sometimes, 1-10 is printed, and other times, the full 1-11 is printed.
	 */
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, waitgroup *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
	waitgroup.Done()
}
