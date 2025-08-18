package main

import (
	"fmt"
)

type Car struct {
	Model  string
	Color  string
	Engine engine
}

type engine struct {
	brand string
	ph    int
}

func main() {
	fmt.Println("Hello World")
	
}
