package exercicio

import "fmt"

type PeopleN struct {
	Name                    string
	Lastname                string
	FavoriteIceCreamFlavors []string
}

func MapWithStruct() {

	clients := map[string]PeopleN{
		"Cardoso": {
			Name:                    "Maria",
			Lastname:                "Cardoso",
			FavoriteIceCreamFlavors: []string{"Flocos", "Chocolate", "Morango"}},
		"Soares": {
			Name:                    "Joao",
			Lastname:                "Soares",
			FavoriteIceCreamFlavors: []string{"Chocomenta", "Nozes", "Melancia"}},
	}

	for i, v := range clients {
		fmt.Println(i)
		for _, nv := range v.FavoriteIceCreamFlavors {
			fmt.Println("\t", nv)
		}
		fmt.Println("")
	}

}
