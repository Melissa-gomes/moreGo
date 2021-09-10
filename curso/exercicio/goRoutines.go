package exercicio

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Alguem struct {
	Name     string
	Acordado bool
	Dentes   bool
	Cafe     bool
}

func Ex45() {
	fmt.Println("Aqui começa minha função")

	Melissa := Alguem{"Melissa", true, true, true}
	wg.Add(3)

	go Acordar(Melissa.Acordado)
	go EscovarOsDentes(Melissa.Dentes)
	go Café(Melissa.Cafe)
	wg.Wait()
}

func Acordar(x bool) {
	if !x {
		x = true
		fmt.Println("agora eu estou acordado", x)
	} else {
		fmt.Println("Ja estou acordado eu deveria escovar meus dentes", x)
	}
	wg.Done()
}

func EscovarOsDentes(y bool) {
	if !y {
		y = true
		fmt.Println("Agora vou escovar meus dentes", y)
	} else {
		fmt.Println("Já estou com os dentes escovados, eu deveria ir tomar café da manhã", y)
	}
	wg.Done()
}

func Café(z bool) {
	if !z {
		z = true
		fmt.Println("Vou tomar café antes de trabalhar para ter energia para o dia.", z)
	} else {
		fmt.Println("eu já tomei café, ja posso trabalhar.", z)
	}
	wg.Done()
}
