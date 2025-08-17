package main

import "fmt"

// array

func testArray1() {
	var myArr [4]string
	myArr[0] = "Rustambek"
	myArr[1] = "O'zbekiston"
	myArr[2] = "O'zbekiston"
	myArr[3] = "O'zbekiston"
	fmt.Println(myArr)
}

func testArray2() {
	myArr := [3]int{1, 2, 3}

	fmt.Println(myArr)
}

func testArray3() {
	myArr := [...]int{12, 3, 4, 5, 6, 9}

	myArr2 := [6]int{12, 3, 4, 5, 6, 9}
	fmt.Println(myArr == myArr2)
}

func testArray4() {
	myarr := [3][3]string{{"C#", "GO", "TS"}, {"Java", "Python", "C++"}, {"C", "C++", "C#"}}

	fmt.Println(myarr[0][1])
}

func testArray5() {
	myarr1 := [3]int{12, 3, 4}
	// to'liq nusxa olish bunda xotiradan boshqa joy ochiladi
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)
	fmt.Println(myarr1 == myarr2)

	myarr1[1] = 100
	fmt.Println(myarr1)
	fmt.Println(myarr2)
}

func testArray6() {
	// reference olish
	myarr := [2]int{1, 2}

	myarr2 := &myarr
	fmt.Println(myarr2)

	myarr[0] = 100
	fmt.Println(myarr2)
}
