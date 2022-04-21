package tasks

import (
	"fmt"
	"strings"
)

func ReverseWords() {
	var s string = "Manav is a passionate software developer"
	arr := strings.Fields(s)
	// res := ""
	n := len(arr)
	for i := 0; i < n/2; i++ {
		// fmt.Println(arr[i])
		// res += arr[i]
		// if i != 0 {
		// 	res += " "
		// }
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	fmt.Println(strings.Join(arr, " "))

	fmt.Println(arr)
}

//todo - use select to get stream of data

// You can edit this code!
// Click here and start typing.
// package main

// import "fmt"

func producer(chnl chan rune, str string, str2 string) {
	go func() {
		for _, v := range []rune(str) {
			chnl <- v
		}

	}()

	go func() {
		for _, v := range []rune(str2) {
			chnl <- v
		}
	}()

	defer close(chnl)
	// send each character of a string to channel

}
func main() {
	ch := make(chan rune)
	str := "Manav"
	str2 := "Sinha"
	producer(ch, str, str2)
	for v := range ch {
		fmt.Printf("%c", v)
	}
	//receive the characters of a string from channel
}
