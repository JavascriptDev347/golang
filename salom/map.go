package main

import "fmt"

func Map() {

	statuses := make(map[string]int)

	// mapga element qo'shish
	statuses["pending"] = 1
	statuses["completed"] = 2
	statuses["canceled"] = 3
	statuses["success"] = 4
	fmt.Println(statuses)

	// mapdan qiymat olish
	var success = statuses["success"]
	fmt.Println(success)

	// mapdan qiymat uchirish
	delete(statuses, "success")
	fmt.Println(statuses)
}