package main

import "fmt"

func structLesson() {
	var cars [2]Car

	cars[0] = Car{"BMW", "Black", engine{"NMW", 6}}
	cars[1] = Car{"Mercedes", "White", engine{"Mercedes", 6}}
	fmt.Println(cars)

	carSlice := []Car{{"BMW", "Black", engine{"BMW", 6}}, {"Mercedes", "White", engine{"Mercedes", 6}}}
	fmt.Println(carSlice)
}
