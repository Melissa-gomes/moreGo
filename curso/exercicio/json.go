package exercicio

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func JsonTranformate() {
	u1 := user{"James", 32}
	u2 := user{"MoneyPenny", 27}
	u3 := user{"M", 54}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sb, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	user1, err1 := json.Marshal(u1)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(string(user1))

	user2, err2 := json.Marshal(u2)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(user2))

	user3, err3 := json.Marshal(u3)

	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(string(user3))

}
