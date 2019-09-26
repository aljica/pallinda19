package main

import (
  "fmt"
)

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan <- int) {
    // TODO
    total := 0
    for i := 0; i < len(a); i++ {
      total += a[i]
    }
    res <- total
}

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7}
    n := len(a)
    ch := make(chan int)
    go Add(a[:n/2], ch)
    go Add(a[n/2:], ch)

    // TODO: Get the subtotals from the channel and print their sum.

    totalSum := 0
    for i:=0; i<2; i++ {
      totalSum += <-ch
    }
    fmt.Println(totalSum)

    // what about for val := range ch {} ??? check it out...?
}
