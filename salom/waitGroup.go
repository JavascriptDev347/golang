package main

import (
	"fmt"
	"sync"
	"time"
)

func waitGroup() {
	var wg sync.WaitGroup
	wg.Add(3)
	go display0("Salom", &wg)
	go display0("Hello", &wg)

	go func() {
		display3("Alhamdulillah ")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")
}

// birinchi usul waitGroup
func display0(input string, wg *sync.WaitGroup) {
	for i := 1; i < 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

// ikkinchi usuli?
func display3(input string) {
	for i := 1; i < 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
