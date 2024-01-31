package channels

import "fmt"

func Example3() {
	ch := make(chan int)
	go publisher(ch)
	reader(ch)
}

func reader(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
