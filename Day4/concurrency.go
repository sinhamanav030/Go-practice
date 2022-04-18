package Day4

import (
	"encoding/json"
	"fmt"
	"sync"

	// "strings"
	"time"
)

func concur() {
	fmt.Println("Hello from concur")
}

func numbers() {
	for i := 1; i < 100; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func alpahabets() {
	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}

func Routine() {
	// go concur()
	// fmt.Println("Hello from Routine")

	go numbers()
	go alpahabets()
	time.Sleep(3000 * time.Millisecond)
}

func hello(done chan bool) {
	fmt.Println("Hello from hello")
	time.Sleep(1000 * time.Millisecond)
	done <- true
}

func getDigits(num int, dig chan int) {
	for num != 0 {
		dig <- num % 10
		num /= 10
	}
	close(dig)
}

func calSquares(num int, sqr chan int) {
	res := 0
	dig := make(chan int)
	go getDigits(num, dig)
	for rem := range dig {
		res += (rem * rem)
	}
	sqr <- res
}

func calCubes(num int, cub chan int) {
	res := 0
	dig := make(chan int)
	go getDigits(num, dig)
	for rem := range dig {
		res += (rem * rem * rem)
	}
	cub <- res
}

func sendChanTut(sendch chan<- int) {
	sendch <- 10
}

func producer(chl chan int) {
	for i := 1; i < 10; i++ {
		chl <- i
	}
	close(chl)
}

func ChannelTut() {
	var a chan int
	if a == nil {
		fmt.Println("channel is nil")
		a = make(chan int)
		fmt.Printf("%T\n", a)
	}

	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("Exececution Success")
	sqr, cub := make(chan int), make(chan int)

	go calCubes(123, cub)
	go calSquares(123, sqr)
	squres, cubes := <-sqr, <-cub
	fmt.Println(squres, cubes)

	sendch := make(chan int)
	go sendChanTut(sendch)
	res := <-sendch
	fmt.Println(res)

	chnl := make(chan int)

	go producer(chnl)
	// for {
	// 	v, ok := <-chnl
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Println(v, ok)
	// }

	for v := range chnl {
		fmt.Println(v)
	}
}

func write(a chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("successfully wrote", i, "to ch")
		a <- i
	}
	// a <- 0
	// a <- 2
	c := 10
	fmt.Println("c", c)
	fmt.Println("here we run")

	close(a)
}

func getVal(a chan int) {
	// for v := range a {
	// 	fmt.Println(v)
	// 	fmt.Println("here in getVal")
	// }
	fmt.Println("1 value->", <-a)
	fmt.Println("2 value->", <-a)
	fmt.Println("3 value->", <-a)
	fmt.Println("4 value->", <-a)
	fmt.Println("5 value->", <-a)
	fmt.Println("still here in getVal")

}

func BufferedChannelTut() {

	// a <- 1
	// fmt.Println(<-a)
	a := make(chan int)
	go write(a)
	// // time.Sleep(2 * time.Second)
	for v := range a {

		fmt.Println("here main")
		fmt.Println(v, cap(a))
		// time.Sleep(2 * time.Second)
	}
	// fmt.Println("Aftr")
	// b := make(chan int)
	// go getVal(b)
	// fmt.Println("wrote 1")
	// b <- 1
	// fmt.Println("wrote 2")
	// b <- 2
	// // close(b)
	// // time.Sleep(time.Second)
	// fmt.Println("wrote 3")
	// b <- 3
	// fmt.Println("wrote 4")
	// b <- 4
	// fmt.Println("wrote 5")
	// b <- 5
	// fmt.Println("wrote 6")
	// // b <- 6
	// close(b)
}

type person struct {
	Month string
	Day   int
}
type response1 struct {
	Page int
}

func EncodingTut() {
	p := &person{
		Month: "May",
		Day:   26,
	}
	fmt.Println(p)
	p_enc, _ := json.Marshal(p)
	fmt.Println(string(p_enc))
	res1D := &response1{
		Page: 1,
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
}

func process(wg *sync.WaitGroup) {
	fmt.Println("Hello in process")
	time.Sleep(time.Second)
	wg.Done()
}
func WaitGroupTut() {
	var wg sync.WaitGroup
	wg.Add(1)
	go process(&wg)
	wg.Wait()
	fmt.Println("done")
}
