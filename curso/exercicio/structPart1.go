package exercicio

import "fmt"

type People struct {
	Name                    string
	Lastname                string
	FavoriteIceCreamFlavors []string
}

func IniciandoStruct() {
	Maria := People{
		Name:                    "Maria",
		Lastname:                "Cardoso",
		FavoriteIceCreamFlavors: []string{"Flocos", "Chocolate", "Morango"},
	}

	Joao := People{
		Name:                    "Joao",
		Lastname:                "Soares",
		FavoriteIceCreamFlavors: []string{"Chocomenta", "Nozes", "Melancia"},
	}

	for _, sabores := range Maria.FavoriteIceCreamFlavors {
		fmt.Println(sabores)
	}

	fmt.Println("")

	for _, sabores := range Joao.FavoriteIceCreamFlavors {
		fmt.Println(sabores)
	}
}
