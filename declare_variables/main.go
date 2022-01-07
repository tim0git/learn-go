package main

import "fmt"

func main() {
	const string1 = "hello world"
	const number1 = 12345
	const boolean1 = true

	var string2 string = "hello world"
	var number2 int = 12345
	var boolean2 bool = true

	var array2 [2]string
	array2[0] = "hello"
	array2[1] = "world"

	string3 := "hello world"
	number3 := 12345
	boolean3 := true
	array3 := [6]int{2, 3, 5, 7, 11, 13}

	type area struct {
		width  int
		height int
	}

	object3 := area{width: 10, height: 5}

	var string4 = "hello world"
	var number4 = 12345
	var boolean4 = true

	fmt.Println(string4)
	fmt.Println(number4)
	fmt.Println(boolean4)

	fmt.Println(string3)
	fmt.Println(number3)
	fmt.Println(boolean3)
	fmt.Println(array3)
	fmt.Println(object3)
	fmt.Println(object3.width)

	fmt.Println(string2)
	fmt.Println(number2)
	fmt.Println(boolean2)
	fmt.Println(array2)

	fmt.Println(string1)
	fmt.Println(number1)
	fmt.Println(boolean1)
}
