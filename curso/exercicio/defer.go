package exercicio

import "fmt"

func UsingDefer() {
	defer fmt.Println("esse eu escrevi primeiro, porem vai aparecer por ultimo")
	fmt.Println("esse foi o segundo que eu escrevi mas vai aparecer primeiro")
	fmt.Println("esse foi o ultimo que eu escrevi mas vai aparecer em segundo lugar")
}
