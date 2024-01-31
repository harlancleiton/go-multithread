package channels

func Example5() {
	hello := make(chan string)
	go receives("Harlan", hello)
	read(hello)
}

func receives(name string, hello chan<- string) {
	hello <- name
}

func read(name <-chan string) {
	println(<-name)
}
