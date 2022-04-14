package Day3

import (
	"fmt"
)

type myInt int

func (a myInt) add(b myInt) {
	fmt.Println(a + b)
}

type Address struct {
	city, state string
}

type Employee struct {
	name     string
	salary   int
	currency string
	Address
}

func (e Employee) displaySalary() {
	e.salary = 1000
	fmt.Println(e.salary)
}

func (e *Employee) changeCurrency(currency string) {
	e.currency = currency
}

func MethodsTut() {

	emp := Employee{
		name:     "Manav",
		salary:   1222,
		currency: "Dollar",
		Address: Address{
			city:  "Delhi",
			state: "Uttar Pradesh",
		},
	}
	emp.displaySalary()
	fmt.Println(emp.name, emp.city, emp.state, emp.currency)
	// fmt.Println(rectangle{10, 20}.Area())
	// fmt.Println(circle{2}.Area())

	emp.changeCurrency("INR")
	(&emp).displaySalary()
	fmt.Println(emp.currency, emp.salary)

	var a myInt = 3
	a.add(4)

}

type shape interface {
	calculateArea
	calculatePerimeter
}

type calculateArea interface {
	Area() float32
}

type calculatePerimeter interface {
	Perimeter() float32
}

type circle struct {
	radius int
}

type rectangle struct {
	width, length int
}

func (c *circle) Area() float32 {
	return 3.14 * float32(c.radius*c.radius)
}

func (r *rectangle) Perimeter() float32 {
	return 2 * float32(r.length+r.width)
}

func (r *rectangle) Area() float32 {
	return float32(r.length * r.width)
}

func (c *circle) Perimeter() float32 {
	return 2 * 3.14 * float32(c.radius)
}

func describe(i interface{}) {
	fmt.Printf("%T\n", i)
}

func InterfacesTut() {
	r1 := rectangle{10, 20}
	c1 := circle{10}
	shapes := []shape{&r1, &c1}
	// (&r1).Area()

	for _, shape := range shapes {
		fmt.Printf("%f %T\n", shape.Area(), shape)
	}

	describe(3)
	describe("aaa")
	describe(r1)

	var i1 shape = &c1
	fmt.Printf("%T\n", i1)

	var i2 shape = &c1
	fmt.Printf("%T\n", i2)

	var i3 shape
	fmt.Println(i3)

}
