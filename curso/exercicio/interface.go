package exercicio

import (
	"fmt"
	"math"
)

type Quadrado struct {
	ladoA float64
}

func (x Quadrado) area() {
	fmt.Println(x.ladoA * x.ladoA)
}

type Circulo struct {
	raio float64
}

func (x Circulo) area() {
	fmt.Println(2 * math.Pi * x.raio)
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func TreinoDeInterface() {
	quadrado := Quadrado{ladoA: 12.}
	circulo := Circulo{raio: 22.}

	medida(quadrado)
	medida(circulo)
}
