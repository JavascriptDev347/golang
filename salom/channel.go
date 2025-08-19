package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Chanel() {
	// channel create
	// channel := make(chan string, 2) //buffer qilish size beriladi
	channel := make(chan string)
	// channelga malumot berish
	// channel <- "Rustambek"
	// 	channel <- "O'zbekiston"
	go func() {
		channel <- "Rustambek"
		channel <- "O'zbekiston"
	}()
	// channeldan malumot olish
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	channel2 := make(chan int)
	go getRandom(channel2)
	for number := range channel2 {
		fmt.Println(number)
	}
}

func getRandom(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 3; i++ {
		number := rand.Intn(1000)
		time.Sleep(1 * time.Second)
		channel <- number
	}
	close(channel)
}
