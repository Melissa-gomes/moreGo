package exercicio

import "fmt"

var d int = 42
var e string = "James Bond"
var f bool = true

func Ex3() {
	s := fmt.Sprintf("%v %v %v\n", d, e, f)
	fmt.Print(s)
}
