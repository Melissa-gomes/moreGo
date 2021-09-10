package exercicio

import "fmt"

func Channels() {
	q := make(chan int)
	c := gene(q)

	receiveT(c, q)

	fmt.Println("about to exit")
}

func gene(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receiveT(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
