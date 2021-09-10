package exercicio

import "fmt"

func Returnsfunc() {
	fmt.Println(returnInt())
	fmt.Println(intAndString())
}

func returnInt() int {
	x := 4
	y := 6
	return x + y
}

func intAndString() (int, string) {
	x := 10
	y := 5

	result := x - y

	return result, "é o resultado de x - y"
}
