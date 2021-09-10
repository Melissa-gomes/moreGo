package exercicio

import "fmt"

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return "deu zica"
}

func erroComParametro(e error) {
	fmt.Print("Passaram esse erro ae", e)
}

func NewErr() {
	outroerr := erroEspecial{"aaaa"}

	erroComParametro(outroerr)
}
