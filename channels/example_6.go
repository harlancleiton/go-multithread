package channels

import (
	"fmt"
	"time"
)

func Example6() {
	data := make(chan int)
	workersAmount := 10
	for i := 0; i < workersAmount; i++ {
		go worker(i+1, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for i := range data {
		fmt.Printf("Worker %d received %d\n", workerId, i)
		time.Sleep(time.Second)
	}
}
