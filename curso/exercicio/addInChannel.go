package exercicio

import "fmt"

func ForWithChannel() {
	channel := make(chan int)
	go addInChannel(channel)
	for v := range channel {
		fmt.Println("--", v)
	}
}

func addInChannel(c chan int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}
