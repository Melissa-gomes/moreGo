package exercicio

import "fmt"

func SliceinSlice() {
	slice := [][]string{
		{
			"Melissa", "Gomes", "Ler",
		},
		{
			"Vitor", "Rodrigues", "cozinhar",
		},
		{
			"Lucas", "Zago", "RPG",
		},
	}

	for _, values := range slice {
		fmt.Println(values[0])
		for _, item := range values {
			fmt.Println("\t", item)
		}
	}
}
