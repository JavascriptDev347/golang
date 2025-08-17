package main

import "fmt"

func types() {
	// Integer
	var X uint8 = 10
	fmt.Println(X+1, X)
	var Y uint16 = 20
	fmt.Println(Y+1, Y)

	// float
	a := 3.23
	b := 4.78
	fmt.Println(a + b)

	// complex
	var d complex128 = complex(6, 2)
	var e complex64 = complex(9, 2)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Printf("D type %T va E type %T", d, e)

	// boolean
	str1 := "Virtual dars"
	str2 := "O'zbekiston"
	str3 := "O'zbekiston"
	result1 := str1 == str2
	result2 := str2 == str3
	fmt.Println(result1, result2)
	fmt.Printf("result1 type %T va result2 type %T", result1, result2)

	// String
	str := "Rustambek"
	fmt.Printf("\nstr uzunligi", len(str))
	fmt.Printf("\nstr Type %T", str)

}
