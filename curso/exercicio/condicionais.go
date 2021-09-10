package exercicio

import "fmt"

func Condicionais() {
	a := 5
	b := 10
	c := 5

	r := a == c
	fmt.Println(r)
	fmt.Println("---------")

	r = a <= b
	fmt.Println(r)
	fmt.Println("---------")

	r = a >= b
	fmt.Println(r)
	fmt.Println("---------")

	r = a < b
	fmt.Println(r)
	fmt.Println("---------")

	r = a > b
	fmt.Println(r)
	fmt.Println("---------")

	r = a != b
	fmt.Println(r)
	fmt.Println("---------")
}
