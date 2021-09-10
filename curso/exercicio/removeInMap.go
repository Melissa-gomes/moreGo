package exercicio

import "fmt"

func RemoveInMap() {
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
	// remove um elemento a um map
	delete(ob, "Sordi_Bruno")

	for k, v := range ob {
		fmt.Println(k)
		for _, va := range v {
			fmt.Println("\t", va)
		}
	}
}
