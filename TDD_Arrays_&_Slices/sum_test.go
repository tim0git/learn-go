package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("should return sum when passed an array of fixed length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		recieved := Sum(numbers)

		expected := 15

		if recieved != expected {
			t.Errorf("expected '%d' but got '%d'", expected, recieved)
		}
	})

	t.Run("should return sum when passed an slice of any length", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		recieved := Sum(numbers)

		expected := 6

		if recieved != expected {
			t.Errorf("expected '%d' but got '%d'", expected, recieved)
		}
	})
}

func BenchmarkSum(b *testing.B) {

	for i := 0; i < b.N; i++ {

		numbers := []int{1, 2, 3, 4, 5}

		Sum(numbers)
		// 15
	}
}

func ExampleSum() {

	numbers := []int{1, 2, 3, 4, 5}

	output := Sum(numbers)

	fmt.Println(output)

	// Output: 15
}

func TestSumAll(t *testing.T) {

	t.Run("should ", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		moreNumbers := []int{1, 2, 3, 4, 5, 6}

		recieved := SumAll(numbers, moreNumbers)

		expected := []int{6, 21}

		if !reflect.DeepEqual(recieved, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, recieved)
		}
	})
}

func BenchmarkSumAll(b *testing.B) {

	for i := 0; i < b.N; i++ {

		numbers := []int{1, 2, 3, 4, 5}

		moreNumbers := []int{1, 2, 3}

		SumAll(numbers, moreNumbers)
		// [15, 6]
	}
}

func ExampleSumAll() {

	numbers := []int{1, 2, 3, 4, 5}

	moreNumbers := []int{1, 2, 3}

	output := SumAll(numbers, moreNumbers)

	fmt.Println(output)

	// [15, 6]
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, recieved, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(recieved, expected) {
			t.Errorf("got %v want %v", recieved, expected)
		}
	}

	t.Run("when passed an array with two characters returns the sum of the tails", func(t *testing.T) {
		recieved := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, recieved, expected)

	})

	t.Run("when passed an array with one character returns the sum of the tails", func(t *testing.T) {
		recieved := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}

		checkSums(t, recieved, expected)

	})
}

func BenchmarkSumAllTails(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 9}, []int{0, 9})
		// [9, 9]
	}
}

func ExampleSumAllTails() {

	output := SumAllTails([]int{1, 9}, []int{0, 9})

	fmt.Println(output)

	// [9, 9]
}
