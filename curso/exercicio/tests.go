package exercicio

import "fmt"

func Testando() {
	umaSoma := somaPraMim(5, 5)

	fmt.Println(umaSoma)
}

func somaPraMim(x, y int) int {
	return x + y
}
