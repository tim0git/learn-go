package main

// Sum when provided an array returns the sum of all the int within the array in sequential order
func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll when provided arrays of numbers returns the sum of all the int within the array in sequential order
func SumAll(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails sums all numbers in an array with the exception of the number at index 0
func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
