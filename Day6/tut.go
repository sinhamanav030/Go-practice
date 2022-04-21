package Day6

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Incrementor(s string) chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func Puller(s string, c chan int) chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for v := range c {
			fmt.Println(s, v)
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

func ConcucrTut() {
	c1 := Incrementor("Foo:")
	c2 := Incrementor("Bar:")
	c3 := Puller("g1:", c1)
	c4 := Puller("g2:", c2)
	fmt.Println(<-c3, <-c4)
	fact := Factorial(5)
	fmt.Println(<-fact)
}

func Factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := 1; i <= n; i++ {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func PipeLineTut() {
	for n := range sq(sq(gen(2, 3, 4))) {
		fmt.Println(n)
	}

}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func FanInTut() {
	c := FanIn(boring("Joe"), boring("bar"))
	for i := 0; i <= 10; i++ {
		fmt.Println(<-c)
	}
}

func boring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()
	return c
}

func FanIn(ip1, ip2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-ip1
		}
	}()

	go func() {
		for {
			c <- <-ip2
		}
	}()
	return c
}
func FanInFanoutTut() {
	s := make([]int, 1000)
	for i := 0; i < 10000; i++ {
		s = append(s, rand.Intn(100000))
	}
	c := generate(s...)
	start := time.Now()
	c1 := square(c)
	c2 := square(c)
	// c3 := square(c)

	for v := range merge(c1, c2) {
		fmt.Printf("%d,", v)
	}
	fmt.Println()
	end := time.Now()

	fmt.Println("Time Utilised:", end.Sub(start))
}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func square(ip <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ip {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, c := range chans {
		go func(ch <-chan int) {
			for v := range ch {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
