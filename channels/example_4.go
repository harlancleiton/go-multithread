package channels

import (
	"fmt"
	"sync"
)

func Example4() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publisherEx4(ch)
	go readerEx4(ch, &wg)
	wg.Wait()
}

func readerEx4(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
		wg.Done()
	}
}

func publisherEx4(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
