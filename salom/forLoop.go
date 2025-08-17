package main

import "fmt"

func forLoop() {
	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// infinite loop strconv.Itoa => int to string
	// i := 1
	// for i < 10 {
	// 	fmt.Println("Alhamdulillah " + strconv.Itoa(i))
	// 	i++
	// }

	myarr := [3]string{"Slaom", "Hello", "bonjurehe"}
	mymap := map[int]string{1: "Slaom mapo map", 2: "Hello", 3: "bonjurehe mape"}

	for key, value := range mymap {
		fmt.Println(key, value)
	}
	for index, value := range myarr {
		fmt.Println(index, value)
	}
}
