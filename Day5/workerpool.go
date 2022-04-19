package Day5

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	id, randomNo int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job)
var result = make(chan Result)
var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())

}

func digits(num int) int {
	res := 0
	for num != 0 {
		rem := num % 10
		res += rem
		num /= 10
	}
	time.Sleep(time.Second)
	return res
}

func worker(wg *sync.WaitGroup) {
	// mutex.Lock()
	for job := range jobs {
		output := Result{job, digits(job.randomNo)}
		result <- output
	}
	// mutex.Unlock()
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(result)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func results(done chan<- bool) {
	for res := range result {
		fmt.Printf("job id :%d\trandom number :%d\tsum of digits:%d\n", res.job.id, res.job.randomNo, res.sumOfDigits)
	}
	done <- true
}

func WorkerPoolTut() {
	fmt.Println(runtime.NumCPU(), runtime.NumGoroutine())
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go results(done)
	noOfWorkers := 5
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Difference :", diff)
}

func server1(res chan string) {
	// time.Sleep(time.Second)
	res <- "from server 1"
}

func server2(res chan string) {
	// time.Sleep(2 * time.Second)
	res <- "from server 2"
}

func SelectTut() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go server1(chan1)
	go server2(chan2)

	select {
	case s1 := <-chan1:
		fmt.Println(s1)
	case s2 := <-chan2:
		fmt.Println(s2)
	}

}

/// go1 blocked for reciving in ch

/// go2

//  main send ch
