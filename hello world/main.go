package main

import "fmt"

// func sumDigits(number int) int {
// 	remainder := 0
// 	sumResult := 0
// 	for number != 0 {
// 		remainder = number % 10
// 		sumResult += remainder
// 		number = number / 10
// 	}
// 	return sumResult
// }

// func main() {

// 	sum := 0

// 	for i := 10; i < 10000; i++ {

// 		sumDig := sumDigits(i)

// 		sum += sumDig - i

// 		if sum < 200 && sum > -200 {
// 			fmt.Println(i)
// 			fmt.Println(sum)
// 		}

// 	}
// }

func main() {

	strings := []string{"one", "two", "three", "four", "five"}

	for i, _ := range strings {
		fmt.Println(i)
		// 1 2 3 4 5

		fmt.Println(fmt.Sprint(65) == "65")

		// "one" "two" "three" "four" "five"
	}
}
