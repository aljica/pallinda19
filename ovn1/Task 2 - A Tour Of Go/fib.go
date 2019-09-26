package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	previ := -1
	return func() (q int) {
		if i == 0 && previ == -1 {
			i++
			return 0
		}
		if i == 1 && previ == -1 {
			previ = 0
			return 1
		}
		q = i + previ
		previ = i
		i = q
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
