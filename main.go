package main

import (
	"fmt"
	"tutorial/Day1"
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
	Day1.Hello()
	Day1.Conditionals(100, 12)
	Day1.Loops(1, 10, 7, 2)
}
