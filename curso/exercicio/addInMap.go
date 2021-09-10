package exercicio

import "fmt"

func Addinmap() {
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
	// adiciona um novo elemento a um map
	ob["Rodrigues_Vitor"] = []string{"cozinhar", "ouvir Lofi", "ver serie"}

	for k, v := range ob {
		fmt.Println(k)
		for _, va := range v {
			fmt.Println("\t", va)
		}
	}
}
