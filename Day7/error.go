package Day7

import (
	"fmt"
	// "runtime"
	"sync"

	// "net"
	"path/filepath"
	// "os"
)

type person struct {
	firstname, lastname string
}

func (p person) display() {
	fmt.Println(p.firstname, p.lastname)
}

func DeferTut() {
	p := person{
		firstname: "Manav",
		lastname:  "Sinha",
	}

	defer p.display()
	fmt.Println("Welcome")

	p.firstname = "Sujit"
	// p.display()

	s := "Manav"

	for _, v := range []rune(s) {
		defer fmt.Printf("%c", v)
	}
}

func ErrorTut() {
	// f, err := os.Open("./test.txt")
	// if err != nil {
	// 	if pErr, ok := err.(*os.PathError); ok {
	// 		fmt.Println(pErr.Path)
	// 	}
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(f.Name())

	// addr, err := net.LookupHost("golangbot123.com")
	// if err != nil {
	// 	if dnsErr, ok := err.(*net.DNSError); ok {
	// 		if dnsErr.Timeout() {
	// 			fmt.Println("operation timed out")
	// 			return
	// 		}
	// 		if dnsErr.Temporary() {
	// 			fmt.Println("temporary error")
	// 			return
	// 		}
	// 		fmt.Println("Generic DNS error", err)
	// 		return
	// 	}
	// 	fmt.Println("Generic error", err)
	// 	return
	// }
	// fmt.Println(addr)

	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)

}

func GoThread() {
	// runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	v := []int{}
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			for j := 0; j < 10; j++ {
				v = append(v, j)
			}
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(v), v)
}
