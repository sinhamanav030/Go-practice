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
