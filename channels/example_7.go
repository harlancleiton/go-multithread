package channels

import "time"

func Example7() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 2
	}()

	select {
	case msg1 := <-ch1:
		println("Received", msg1)
	case msg2 := <-ch2:
		println("Received", msg2)
	case <-time.After(time.Second * 3):
		println("Timeout")
		// default:
		// 	println("Default")
	}
}
