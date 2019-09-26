package main

import (
	"fmt"
	"strconv"
	"time"
)

func Remind(text string, delay time.Duration, c chan<- string) {
	//delay *= time.Second
	delay *= time.Hour
	time.Sleep(delay)

	hr := time.Now().Hour()
	min := time.Now().Minute()
	//sec := time.Now().Second()
	c <- "Klockan är " + strconv.Itoa(hr) + "." + strconv.Itoa(min) /*+ "." + strconv.Itoa(sec)*/ + ": " + text
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	fmt.Println(time.Now())

	go Remind("Dags att äta", 3, c1)
	go Remind("Dags att arbeta", 8, c2)
	go Remind("Dags att sova", 24, c3)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
			go Remind("Dags att äta", 3, c1)
		case msg2 := <-c2:
			fmt.Println(msg2)
			go Remind("Dags att arbeta", 8, c2)
		case msg3 := <-c3:
			fmt.Println(msg3)
			go Remind("Dags att sova", 24, c3)
		}
	}
}
