package main

import (
	"fmt"
	"sort"
)

func testSlice1() {
	// yangi slice yaratish
	myslice := []int{1, 2, 4}

	// slice oxiriga yangi element qushish
	myslice = append(myslice, 5)

	fmt.Println("slice uzunligi", len(myslice))
}

func testSlice2() {
	// sliceni birorta qatordan yaratib olish
	myarr := [7]string{"olma", "anor", "banan", "nok", "qovun", "shaptoli", "tarvuz"}

	myslice := myarr[1:4]

	fmt.Println(myslice)
}

func testSlice3() {
	// slice tartiblash
	myslice := []int{34, 56, 345446, 323, 32343, 3, 345, 3432421, 123, 24, 31, 34, 5, 23, 45}
	sort.Ints(myslice)
	fmt.Println(myslice)
}
