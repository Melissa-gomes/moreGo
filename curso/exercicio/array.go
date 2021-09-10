package exercicio

import "fmt"

func Array() {
	valores := [5]int{1, 2, 3, 4, 5}

	for _, valor := range valores {
		fmt.Println(valor)
	}

	fmt.Printf("\n%T", valores)
}
