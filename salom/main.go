package main

import (
	"fmt"
	"time"
)

// channel => go routinela aloqasi uchun
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Tez"
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			channel2 <- "Sekin"
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		default:
			fmt.Println("Ma'lumot yo'q")
		}
	}
}
