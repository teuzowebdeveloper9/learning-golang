package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var p *int

	fmt.Printf("Zero value of int: %d\n", i)
	fmt.Printf("Zero value of float64: %f\n", f)
	fmt.Printf("Zero value of bool: %t\n", b)
	fmt.Printf("Zero value of string: '%s'\n", s)
	fmt.Printf("Zero value of *int: %v\n", p)
}
