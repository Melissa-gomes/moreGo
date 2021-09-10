package exercicio

import "testing"

func TestSoma(t *testing.T) {
	test := somaPraMim(5, 5)
	resultado := 10

	if test != resultado {
		t.Error("Expected:", resultado, "Got:", test)
	}
}
