package Day2

import (
	"fmt"
	"reflect"
	"strings"
)

func SwitchTut() {
	num := 12
	switch num {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("You are special one")
	}

	char := 'd'
	switch char {
	case 'a', 'i', 'e', 'o', 'u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	switch {
	case num < 10 && num > 5:
		fmt.Println("5-10")
	case num > 10 && num < 20:
		fmt.Println("10-20")
	default:
		fmt.Println(">20 || <5")
	}

	switch {
	case num > 0:
		if num < 0 {
			break
		}
		fmt.Println(">0")
		fallthrough
	case num > 5:
		fmt.Println(">5")
		fallthrough
	case num > 10:
		fmt.Println(">10")
		fallthrough
	default:
		fmt.Println(">20")
	}

}

func VariadicTut(a ...int) {
	for i := range a {
		fmt.Println(a[i])
	}
}

func MapTut() {
	emp := make(map[string]int)
	// emp["manav"] = 1000
	// fmt.Println(emp)

	emp4 := make(map[string]int)
	emp4["manav"] = 1000
	// fmt.Println(emp4)

	fmt.Println(reflect.DeepEqual(emp, emp4))

	emp2 := map[string]int{
		"steve": 100,
		"jamie": 200,
	}
	fmt.Println(emp2)

	var emp3 map[string]int
	// emp3["a"] = 1
	// fmt.Println(emp3)
	fmt.Println(reflect.DeepEqual(emp, emp3))

	value, ok := emp2["manav"]
	fmt.Println(value, ok)

	value, ok = emp2["jamie"]
	fmt.Println(value, ok)

	// for key, value := range emp2 {
	// 	fmt.Println(key, value)
	// }

	delete(emp2, "jamie")
	// for key, value := range emp2 {
	// 	fmt.Println(key, value)
	// }
}

func StringTut() {
	var s string = "Señor"
	fmt.Printf("%s\n", s)
	for i := range s {
		fmt.Printf("%c ", s[i])
	}
	println()
	st := []rune(s)

	snew := strings.Clone(s)

	fmt.Println(snew)

	for i := 0; i < len(st); i++ {
		fmt.Printf("%c ", st[i])
	}

	for index, rune := range s {
		fmt.Printf("%d %c\n", index, rune)
	}

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)

	b2 := []byte{67, 97, 102}
	str2 := string(b2)
	fmt.Println(str2)
}

func Array() {
	fmt.Println("Hello from Array")
	var a [3]int
	fmt.Println(a)

	a2 := [3]int{12, 14, 15}
	fmt.Println(a2)

	a3 := [...]int{11, 11}
	fmt.Println(a3)

	// cannot use a2 (variable of type [3]int) as [2]int value in assignmentcompilerIncompatibleAssign
	// a3=a2

	a4 := a3
	a4[0] = 13
	fmt.Println(a4, len(a4))

	for i, v := range a4 {
		fmt.Println(a4[i], "->", v)
	}

	c := [][]int{
		{1, 2},
		{2, 3},
	}

	fmt.Println(c)

	a11 := [...]int{1, 2, 3, 4, 5, 6, 7}

	a12 := a11[1:4]
	var a13 []int = a11[1:5]
	fmt.Println(a11, a12, a13)

	a14 := [...]int{1, 2, 3, 4, 5}

	a15 := a14

	a15[0] = 3

	fmt.Println(a14, a15)

	a16 := []int{1, 2, 3, 4}

	a17 := a16
	a18 := a16
	// a17[0] = 3
	// a18[1] = 4
	fmt.Println(a16, a17, a18)
	a19 := a16[0:2]
	println(len(a19), cap(a19))

	a20 := a19[:cap(a19)]
	fmt.Println(a19, a20)

	a21 := make([]int, 10, 10)

	fmt.Println(a21)

	a22 := []int{1, 2, 3}
	fmt.Println(a22, cap(a22))

	a22 = append(a22, 4)
	fmt.Println(a22, cap(a22))

	a23 := make([]int, 2, 5)

	fmt.Println(a23, cap(a23))
	println(a23)
	a23 = append(a23, 5)
	fmt.Println(a23, cap(a23))
	println(a23)

	a23 = append(a23, a22...)
	println(a23)
	fmt.Println(a23, cap(a23))

	for i := range a23 {
		fmt.Println(a23[i])
	}
	a24 := make([]int, len(a23))
	copy(a24, a23)

	println(a24, a23)
	fmt.Println(len(a23), cap(a23), len(a23), cap(a24))
}

func PointerTut() {
	b := 255
	var a *int = &b
	fmt.Println(a, *a)

}

func PonterHelp(a *int) {
	*a = 77
}

func PointerHelp() *int {
	i := 101
	return &i
}

type employee struct {
	firstName, lastName string
	age, salary         int
}

type emp struct {
	string
	int
}

func StructTut() {
	emp := employee{
		firstName: "Manav",
		lastName:  "Sinha",
		age:       21,
		salary:    32000,
	}

	fmt.Println(emp)

	emp2 := &employee{"Gaurav", "Kumar", 22, 32000}

	fmt.Println(emp2)

	emp3 := struct {
		Name string
		Age  int
	}{
		Name: "Manav Sinha",
		Age:  21,
	}

	fmt.Println(emp3)

	fmt.Println(emp.firstName, emp.lastName, emp.age, emp.salary)

	fmt.Println(emp2.firstName, emp2.lastName, emp2.age, emp2.salary)
}
