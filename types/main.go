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
		times := 0
		for i := 0; i < 5; i++ {
			times++

			fmt.Printf("Yes, stephane is beathfull! (%d)\n", times)
		}

		fmt.Printf("You said stephane was beautiful %d times\n", times)
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

	inferred := "I am inferred"

	fmt.Printf("infired type is %T\n", inferred)

	AgainInferred := 3.14

	fmt.Printf("AgainInferred type is %T\n", AgainInferred)
}
