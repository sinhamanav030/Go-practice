package main

import (
	"fmt"
	// "tutorial/Day1"
	"tutorial/Day2"
)

var p = 10

func init() {
	fmt.Println("Main Package Initialized")
	if p == 0 {
		fmt.Println("P is not initializd")
	} else {
		fmt.Println("P is initialized with value:", p)
	}

}

func main() {
	fmt.Println("Hello in Tutorial")
	// Day1.Hello()
	// Day1.Conditionals(100, 12)
	// Day1.Loops(1, 10, 7, 2)
	Day2.SwitchTut()
	Day2.Array()

	Day2.VariadicTut(13, 12)

	Day2.MapTut()

	Day2.StringTut()

	Day2.PointerTut()

	a := 10
	Day2.PonterHelp(&a)

	fmt.Println(a)

	fmt.Println(*(Day2.PointerHelp()))

	Day2.StructTut()

}
