package main

import (
	"fmt"
	"strconv"
	"time"
)

func Remind(text string, delay time.Duration) {
	delay *= time.Hour
	for {
		time.Sleep(delay)
		hr := time.Now().Hour()
		min := time.Now().Minute()
		fmt.Println("Klockan är " + strconv.Itoa(hr) + "." + strconv.Itoa(min) + ": " + text)
	}
}

func main() {
	fmt.Println(time.Now())

	go Remind("Dags att äta", 3)
	go Remind("Dags att arbeta", 8)
	go Remind("Dags att sova", 24)

	select {}
}
