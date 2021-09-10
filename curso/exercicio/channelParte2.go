package exercicio

import "fmt"

func ChannelPart2() {
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("--------\n")
	fmt.Printf("cs\t%T\n", cs)
}
