package exercicio

import "fmt"

func Valoresdoprintf() {

	x := 42

	// %d ===> decimal
	// %#x ===> hexadecimal
	// %b ===> binario

	fmt.Printf("%d, %#x, %b", x, x, x)
}
