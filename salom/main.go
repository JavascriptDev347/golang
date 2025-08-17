package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Assalomu alaykum")
	// variables create with var keyword
	var x, z int = 3, 45
	var y int = 4
	fmt.Println(x + y + z)

	// short version create variables
	a, b, c := 3, 4, 5
	fmt.Println(a, b, c)

	// types()

	testIf()
	testIfElse()
	testSwitch()
}

func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("i is 7")
	}
}

func testIfElse() {
	points := 6000

	if points < 5000 {
		fmt.Println("Your points is less than 5000")
	} else if points < 10000 {
		fmt.Println("Your points is more than 1000")
	} else {
		fmt.Println("Your points is another")
	}
}

func testSwitch() {
	weekend := time.Now().Weekday()

	switch weekend {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 0:
		fmt.Println("Sunday")
	default:
		fmt.Println("XATO")
	}
}
