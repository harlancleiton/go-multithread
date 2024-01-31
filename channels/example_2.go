package channels

func Example2() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever
}
