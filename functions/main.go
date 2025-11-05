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

func main() {

	result := sum(3, 5)

	fmt.Println("The sum is:", result)

	numbers := []int{10, 7, 8, 9, 1, 5}

	sortedNumbers := quicksort(numbers)

	fmt.Println("Sorted numbers:", sortedNumbers)
}
