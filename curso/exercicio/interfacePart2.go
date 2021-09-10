package exercicio

import "fmt"

type umaPessoa struct {
	nome  string
	idade int
}

func (p *umaPessoa) falar() {
	fmt.Println(p.nome, ": eu posso falar, não latir, mas falar.")
}

func dizeralgumacoisa(a humanos) {
	a.falar()
}

type humanos interface {
	falar()
}

func APlicandoInterfaces() {
	theo := umaPessoa{"theo", 12}
	theo.falar()
	dizeralgumacoisa(&theo)
}
