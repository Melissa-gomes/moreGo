package exercicio

import "fmt"

func MoreChannel() {
	c := gen()
	receive(c)
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}

}
