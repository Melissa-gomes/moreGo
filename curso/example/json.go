package example

import (
	"encoding/json"
	"fmt"
)

type NewPerson struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
}

func Json() {
	sininho := NewPerson{Nome: "Tinker", Sobrenome: "Bell", Idade: 22, Profissao: "fada artesã"}

	capitaoGancho := NewPerson{"Capitão", "Gancho", 46, "capitão dos piratas"}

	s, err := json.Marshal(sininho)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(s))

	c, err2 := json.Marshal(capitaoGancho)

	if err2 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(c))
}
