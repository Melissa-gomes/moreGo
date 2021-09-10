package exercicio

import "fmt"

func FuncReturnsFunc() {
	x := inseption()

	x()
}

func inseption() func() {

	return func() {
		fmt.Println("OIii pessoa com quem estou falando")
	}
}
