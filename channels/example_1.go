package channels

import "fmt"

func Example1() {

	ch := make(chan string)

	go func() {
		ch <- "Hello World!"
	}()

	msg := <-ch
	fmt.Println(msg)
}
