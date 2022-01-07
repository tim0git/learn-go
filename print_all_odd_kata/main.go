package main

// print all even numbers from 1 - 9
import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i%2 > 0 {
			fmt.Println(i)
		}

	}
}
