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

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	slice := []string{"apple", "banana", "cherry"}

	slice = append(slice, "date")

	slice = append(slice, "elderberry", "fig", "grape")

	for _, fruit := range slice {
		fmt.Println(fruit)
	}

}
