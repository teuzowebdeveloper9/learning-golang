package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("helllo brow!")
	var stephane string = "stephane"
	fmt.Println(strings.ToUpper(stephane))

	var stephaneBeathfull bool = true

	fmt.Printf("Is stephane beathfull? %t\n", stephaneBeathfull)

	var mankey any = "A wild mankey appears!"
	fmt.Println(mankey)

	mankey = 42

	if age, ok := mankey.(int); ok {
		fmt.Printf("the mankey age is %d\n", age)
	} else {
		fmt.Printf("the mankey age is %v\n", mankey)
	}
}
