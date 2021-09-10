package exercicio

import "fmt"

type Vehicle struct {
	Port  int
	Color string
}

type Caminhonete struct {
	Vehicle
	FourWheels bool
}

type Sedan struct {
	Vehicle
	LuxuryModel bool
}

func MoreStruct() {
	caminhonete := Caminhonete{
		Vehicle:    Vehicle{4, "grey"},
		FourWheels: true,
	}

	sedan := Sedan{
		Vehicle:     Vehicle{4, "black"},
		LuxuryModel: false,
	}

	fmt.Println(caminhonete.Color)
	fmt.Println(sedan.Vehicle)
}
