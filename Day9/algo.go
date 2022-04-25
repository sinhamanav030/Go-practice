package Day9

import (
	// "log"
	// "errors"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"runtime/debug"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("norgate math err %v %v %v", n.lat, n.long, n.err)
}

func ErrorTut() {
	_, err := os.Open("abc.txt")
	if err != nil {
		// log.Println(err)
		// log.Fatalln(err)
		panic(err)
	}
}

func Sqrt(f float32) (float64, error) {
	// if f < 0 {
	// 	return 0, errors.New("Number should be reater than 0")
	// }
	// for i := 1; float32(i*i) <= f; i++ {
	// 	if float32(i*i) == f {
	// 		return float64(i), nil
	// 	}
	// }
	// return 0, errors.New("Cannot find sqrt")

	if f < 0 {
		nme := fmt.Errorf("Norgate math redux:sqrt of -ve number %v", f)
		return 0, &NorgateMathError{"50.2N", "99.245 W", nme}
	}
	return 42, nil

}

type add func(a int, b int) int

func simple(a func(a int, b int) int) {
	fmt.Println(60, 7)
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		return t + " " + b
	}
	return c
}

type student struct {
	firstName, lastName, grade, country string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func FirstClassFunctionTut() {
	a := func() {
		fmt.Println("I am function a")
	}
	a()
	fmt.Printf("%T\n", a)

	// var c add = func(a int, b int) int {
	// 	return a + b
	// }
	// s := c(5, 7)
	// fmt.Println("Sum", s)

	d := func(a int, b int) int {
		return a + b
	}
	simple(d)

	e := simple2()
	fmt.Println(e(10, 2))

	// f := appendStr()

	// fmt.Println(f("Manav"))

	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func PanicRecoverTut() {
	firstName := "Elon"
	// lastname := "Musk"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

}

type order struct {
	ordId  int
	custId int
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func ReflectionTut() {
	o := order{456, 7}
	createQuery((o))
}

func FileHandTut() {

	fptr := flag.String("fpath", "test.txt", "file to read path from")
	flag.Parse()
	fmt.Println(*fptr)
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading err", err)
	} else {
		fmt.Println(string(data))
	}
}

func BuffIoTut() {
	fptr := flag.String("fp", "", "file tor read path from")
	flag.Parse()
	fmt.Println(*fptr)
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	if err = f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	// r := bufio.NewReader(f)
	// b := make([]byte, 3)
	// for {
	// 	n, err := r.Read(b)
	// 	if err != nil {

	// 		fmt.Println("Error reading file:", err)
	// 		break
	// 	}
	// 	fmt.Println(string(b[0:n]))
	// }
	s := bufio.NewScanner(f)
	for s.Scan() {

		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func FileWriteTut() {
	f, err := os.Create("./Day9/wr.txt")
	if err != nil {
		log.Fatal(err)
	}
	l, err := f.WriteString("Hello Manav")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		panic(err)
	}

}
