package main

import (
	"fmt"
	"time"
)

// go display("Salom ")
// 	go display("Hello ")

// 	fmt.Scanln()

func display(input string) {
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}

func display2(input string) {
	fmt.Println(input)
}
