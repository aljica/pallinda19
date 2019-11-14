package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string, 1) // Append 1. Buffered channel.
	/* If we do not buffer the channel to expect specifically ONE item,
	then the line {ch <- "Hello world!"} will wait forever, because
	the channel will expect more data to flow through it, so the program
	will not progress to the next line of code.*/
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
