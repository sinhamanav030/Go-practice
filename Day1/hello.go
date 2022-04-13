package Day1

import (
	"fmt"
	"unsafe"
)

func init() {
	fmt.Println("Day 1 package is initialized")
}

func Loops(initial, final, breaker, cont int) {
outer:
	for i := initial; i <= final; i++ {
		if i == breaker {
			break outer
		}

		if i == cont {
			continue
		}
		fmt.Println(i)
	}

}

func Conditionals(num1, num2 int) {
	if num1 > num2 {
		fmt.Printf("Hello %d is greater than %d\n", num1, num2)
	} else if num1 == num2 {
		fmt.Printf("%d is equal to %d\n", num1, num2)
	} else {
		fmt.Printf("Hello %d is lower than %d\n", num1, num2)
	}

	if num3 := num1 + num2; num3 > 100 {
		fmt.Printf("%d is greater than 100\n", num3)
	} else {
		fmt.Printf("%d is less than 100\n", num3)
	}
}

func variables() {
	var age int
	fmt.Println("Int without intitalization :", age)

	age = 29
	fmt.Println("Int with assignment :", age)

	var val int = 29
	fmt.Println("Int with initialization :", val)

	var temp = 29
	fmt.Println("Variable without type declaration :", temp)

	// cannot use "abcd" (untyped string constant) as int value in assignmentcompilerIncompatibleAssign
	// temp = "abcd"

	//Multiple variable Declaration
	var height, width = 100, 50
	fmt.Println("Multiple Declaration :", height, width)

	//declaring variable of different types
	var (
		myName   = "Manav"
		myAge    = 21
		myHeight int
	)
	fmt.Println("Multiple Declaration :", myName, myAge, myHeight)

	//Short Hand Assignment
	count := 10
	fmt.Println("Short hand Assignment", count)

	//Short HandMultiple Declaration
	newName, newAge := "Manav", 21
	fmt.Println("Short Hand Multiple Declaration", newName, newAge)

	//Error - no new variables on left side of :=compilerNoNewVar
	//  newName,newAge := "Karan",25

	//Short Hand Multiple Assignment
	newName, newAge = "Karan", 25
	fmt.Println("Short Hand Multiple Assignment changed already initaliazed value", newName, newAge)

}

func variableTypes() {
	// The following are the basic types available in Go

	// bool
	// Numeric Types
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	// float32, float64
	// complex64, complex128
	// byte
	// rune
	// string

	var a, b bool = true, false
	fmt.Println("Value of a and b :", a, b)

	bool1, bool2 := true, false
	fmt.Println("Value of first and second bool:", bool1, bool2)

	boolOr := bool1 && bool2
	fmt.Println(boolOr)

	var a8 int8 = 100
	fmt.Printf("Int 8 type variable : %d ,Type is %T ans size is %d\n", a8, a8, unsafe.Sizeof(a8))
	//cannot use 255 (untyped int constant) as int8 value in assignment (overflows)compilerNumericOverflow
	// range: -128 to 127
	// a8 = 255

	// range: -32768 to 32767
	var a16 int16 = 255
	fmt.Printf("Int16 type of variable  :%d ,Type is %T ans size is %d\n", a16, a16, unsafe.Sizeof(a16))

	//range: -2147483648 to 2147483647
	var a32 int32 = 222222222
	fmt.Printf("Int32 type of variable : %d ,Type is %T ans size is %d\n", a32, a32, unsafe.Sizeof(a32))

	// range: -9223372036854775808 to 9223372036854775807
	var a64 int64 = -3333333333333333333
	fmt.Printf("Int64 type of variable : %d ,Type is %T ans size is %d\n", a64, a64, unsafe.Sizeof(a64))

	// int: represents 32 or 64 bit integers depending on the underlying platform.
	// You should generally be using int to represent integers unless there is a need to use a specific sized integer.

	var (
		au8  uint8  = 8
		au16 uint16 = 11111
		au32 uint32 = 3333333333
		au64 uint64 = 3333333333333333333
	)

	fmt.Printf("Value:%d ,Type :%T, size :%d\n", au8, au8, unsafe.Sizeof(au8))
	fmt.Printf("Value:%d ,Type :%T, size :%d\n", au16, au16, unsafe.Sizeof(au16))
	fmt.Printf("Value:%d ,Type :%T, size :%d\n", au32, au32, unsafe.Sizeof(au32))
	fmt.Printf("Value:%d ,Type :%T, size :%d\n", au64, au64, unsafe.Sizeof(au64))

	var (
		af32 float32 = 2222.2
		af64 float64 = 22222222222222.22
	)

	fmt.Printf("Value: %f ,Type :%T, size :%d\n", af32, af32, unsafe.Sizeof(af32))
	fmt.Printf("Value: %f ,Type :%T, size :%d\n", af64, af64, unsafe.Sizeof(af64))

	//Complex type
	c1 := complex(5, 7)
	fmt.Println("Complex Number", c1)

	c2 := 10 + 15i

	// TODO
	//Question - How to access real and imaginary part of complex number
	// c := complex(23, 31)
	// realPart := real(c)    // gets real part
	// imagPart := imag(c)    // gets imaginary part
	fmt.Println("Complex Number with shorthand intialization :", c2)

	//iota is used for succesive number initalization
	const (
		ai = 2016 + iota
		bi
		ci
		di
	)

	fmt.Println(ai, bi, ci, di)

	//String
	var firstName, lastName string = "Manav", "Sinha"
	fmt.Println(firstName + " " + lastName)
}

func typeConversion() {

	aint := 5
	bint := 10.0

	// invalid operation: aint + bint (mismatched types int and float64)compilerMismatchedTypes
	// res := aint+ bint

	res := aint + int(bint)
	fmt.Println(res)

	var c1 = complex(5, 10)

	fmt.Println(complex128(c1))

	var st = 'a'

	fmt.Println(int(st))
}

func constTut() {
	const a = 10
	fmt.Printf("%T\n", a)

	// cannot assign to a (untyped int constant 10)compilerUnassignableOperand
	// a = 66

	const b float32 = a
	fmt.Println(b)

	//create alias of type
	type s string

	const temp s = "12"
	fmt.Println(temp)
	// cannot use temp (constant "12" of type s) as string value in constant declarationcompilerIncompatibleAssign
	// const temp2 string = temp

	var div = 10.0 / 3
	fmt.Println(div)

	var num1, num2 = 10.0, 3

	//invalid operation: num1 / num2 (mismatched types float64 and int)compilerMismatchedTypes
	// var ans = num1/num2

	var ans = num1 / float64(num2)

	fmt.Println(ans)
}

func findArea(radius float32) float32 {
	return 3.14 * radius * radius
}

func findPerimeter(width, height float32) float32 {
	return 2 * (width + height)
}

func findAreaAndPerimeter(width, height float32) (float32, float32) {
	return 2 * (width + height), width * height
}

func findRect(width, height float32) (area, perimeter float32) {
	perimeter = 2 * (width + height)
	area = width * height
	return
}

func Hello() {
	fmt.Println("Hello World")

	variables()

	variableTypes()

	typeConversion()

	constTut()

	fmt.Println(findArea(3.1))

	fmt.Println(findPerimeter(10.4, 14.5))

	var area, rect = findAreaAndPerimeter(5, 10)
	fmt.Println(area, rect)

	area, rect = findRect(5, 10)
	fmt.Println(area, rect)

	//blank identifier _ is used iplace of any value of any type

	area, _ = findAreaAndPerimeter(10, 20)
	fmt.Println(area)
}
