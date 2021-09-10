package exercicio

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func ErrPersonalizado() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errNovo := fmt.Errorf("mano deu algum ruim nas coordenadas que vc me passou irmão")
		return 0, sqrtError{"123", "345", errNovo}
	}
	return 42, nil
}
