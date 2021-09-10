package exercicio

import "fmt"

func MapInGo() {
	// map é o objetoJs do Golang

	ob := map[string][]string{
		"Zago_Lucas": {
			"karaoke", "Kart", "RPG",
		},
		"Sordi_Bruno": {
			"jogar", "beber", "codar",
		},
		"Carneiro_Denner": {
			"comer cachorro quente", "ouvir musica", "ser legal",
		},
	}
	for k, v := range ob {
		fmt.Println(k)
		for _, va := range v {
			fmt.Println("\t", va)
		}
	}
}
