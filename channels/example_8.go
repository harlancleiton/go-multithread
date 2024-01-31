package channels

func Example8() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	println(<-ch)
	println(<-ch)
}
