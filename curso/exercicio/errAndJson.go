package exercicio

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string
	Last   string
	Movies []string
}

func MoreERR() {
	p1 := person{
		"Harry",
		"Potter",
		[]string{"Camara secreta", "Pedra filosofal", "O enigma do principe"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
