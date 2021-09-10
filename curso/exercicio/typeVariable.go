package exercicio

import "fmt"

type coxinha int

var x coxinha

var y int

func Typevariable() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
