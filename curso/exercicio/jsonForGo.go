package exercicio

import (
	"encoding/json"
	"fmt"
)

type jsontogo []struct {
	First  string   `json: "First"`
	Last   string   `json: "Last"`
	Age    int      `json: "Age"`
	Movies []string `json: "Movies"`
}

func JsonForGO() {
	stringInJson := `[{"First":"James","Last":"Bond","Age":32,"Movies":["James bond 1", "James bond 2", "James bond 3"]},{"First":"Tinker","Last":"Bell","Age":16,"Movies":["SEgredo das fadas", "Aquele da irmã", "Aquele do pózinho"]},{"First":"Harry","Last":"Potter","Age":32,"Movies":["Camara secreta", "Pedra filosofal", "Enigma do principe"]}]`

	var result jsontogo

	err := json.Unmarshal([]byte(stringInJson), &result)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
