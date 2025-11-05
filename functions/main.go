package main

import "fmt"

func sum(a int, b int) int {

	return a + b
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func conditional() []bool {
	a := true
	b := false

	if a && !b {
		fmt.Println("Condition met!")
	} else {
		fmt.Println("Condition not met.")
	}

	return []bool{a, b}
}

func subtraction(a int, b int) int {
	if a < b {
		fmt.Printf("We cannot return negative numbers, so we will invert them.")

		return b - a
	}

	return a - b
}

func main() {

	result := sum(3, 5)

	fmt.Println("The sum is:", result)

	numbersSubtracted := subtraction(10, 4)

	fmt.Println("The subtraction is:", numbersSubtracted)

	numbers := []int{10, 7, 8, 9, 1, 5}

	var sortedNumbers []int = quicksort(numbers)

	fmt.Println("Sorted numbers:", sortedNumbers)

	conditional()
}
