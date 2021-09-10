package exercicio

import (
	"encoding/json"
	"fmt"
	"os"
)

type userNew struct {
	First string
	Age   int
}

func ComecandoTratamentodeErr() {
	u1 := userNew{"James", 32}
	u2 := userNew{"MoneyPenny", 27}
	u3 := userNew{"M", 54}

	users := []userNew{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println(err)
	}

}
