package main

import (
	"fmt"
)

func main() {

	var array [5]int

	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50

	intslice := array[:]

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	fmt.Println(intslice)

	slice := []string{"apple", "banana", "cherry"}

	slice = append(slice, "date")

	slice = append(slice, "elderberry", "fig", "grape")

	for _, fruit := range slice {
		fmt.Println(fruit)
	}

	var mixedSlices []any

	for _, n := range intslice {
		mixedSlices = append(mixedSlices, n)
	}

	for _, s := range slice {
		mixedSlices = append(mixedSlices, s)
	}

	fmt.Println(mixedSlices)

}
