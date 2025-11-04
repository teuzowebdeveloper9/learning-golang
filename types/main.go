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

	if stephaneBeathfull {
		fmt.Println("Yes, stephane is beathfull!")
	} else {
		fmt.Println("No, stephane is not beathfull!")
	}

	stephaneBeathfull = false

	if stephaneBeathfull {
		fmt.Println("Yes, stephane is beathfull!")
	} else {
		fmt.Println("No, stephane is not beathfull!")
	}

	var mankey any = "A wild mankey appears!"
	fmt.Println(mankey)

	mankey = 42

	if age, ok := mankey.(int); ok {
		fmt.Printf("the mankey age is %d\n", age)
	} else {
		fmt.Printf("the mankey age is %v\n", mankey)
	}

}
