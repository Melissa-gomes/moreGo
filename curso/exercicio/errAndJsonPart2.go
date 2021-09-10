package exercicio

import (
	"encoding/json"
	"fmt"
	"log"
)

type personN struct {
	First  string
	Last   string
	Movies []string
}

func JsonAndErr() {
	p1 := personN{
		"Harry",
		"Potter",
		[]string{"Camara secreta", "Pedra filosofal", "O enigma do principe"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("CILADA")
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Mano cai nessa não, mo cilada")
	}

	return bs, nil
}
