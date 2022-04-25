package tasks

import (
	"fmt"
	"strings"
	"sync"
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

func producer(chnl chan rune, str1 string, str2 string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _, v := range []rune(str1) {
			chnl <- v
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		for _, v := range []rune(str2) {
			chnl <- v
		}
		defer close(chnl)
	}()

	// send each character of a string to channel

}
func ChannelTask() {
	ch1 := make(chan rune)
	// ch2 := make(chan rune)
	str := "Manav"
	str2 := "Sinha"
	producer(ch1, str, str2)
	for v := range ch1 {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	//receive the characters of a string from channel
}

func producerSelect(chnl1 chan rune, chnl2 chan rune, s1 string, s2 string) {
	for _, v := range []rune(s1) {
		chnl1 <- v
	}
	defer close(chnl1)

	for _, v := range []rune(s2) {
		chnl2 <- v
	}
	defer close(chnl2)
	// q <- true
}

func ChannelTaskSelect() {
	ch1 := make(chan rune)
	ch2 := make(chan rune)
	str := "Manav"
	str2 := "Sinha"
	go producerSelect(ch1, ch2, str, str2)
	// for v := range ch1 {
	// 	fmt.Printf("%c", v)
	// }
	for {
		flag := false
		select {
		// case <-q:
		// 	flag = true
		case v, ok := <-ch1:
			fmt.Printf("%c", v)
			if ok == false {
				fmt.Println()
			}
		case v, ok := <-ch2:
			fmt.Printf("%c", v)
			if ok == false {
				flag = true
			}
		}
		if flag {
			break
		}
	}
	fmt.Println()
}
