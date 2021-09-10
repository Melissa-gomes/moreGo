package exercicio

import "fmt"

func Params() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 14}
	fmt.Println(exercicio1(slice...))
	fmt.Println(exercicio2(slice))
}

func exercicio1(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func exercicio2(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
