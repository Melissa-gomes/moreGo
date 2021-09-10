package exercicio

import "fmt"

func Channel() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	whithBuffer()
}

func whithBuffer() {
	c := make(chan int, 1)

	go func() {
		c <- 28
	}()

	fmt.Println(<-c)
}
