package exercicio

import "fmt"

func SwitchCase() {
	numeroDaSorte := 15

	switch {
	case numeroDaSorte == 7:
		fmt.Printf("%v é meu numero da sorte\n", 7)
	case numeroDaSorte == 15:
		fmt.Printf("%v é meu numero da sorte\n", 15)
	case numeroDaSorte == 18:
		fmt.Printf("%v é meu numero da sorte\n", 18)
	}
}
