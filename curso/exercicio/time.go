package exercicio

import (
	"fmt"
	"time"
)

func Usingtime() {
	anoQueNasci := 2001

	t := time.Now()

	for anoQueNasci <= t.Year() {
		fmt.Println(anoQueNasci)
		anoQueNasci++
	}
}
