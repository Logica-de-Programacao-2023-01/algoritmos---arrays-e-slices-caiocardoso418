package main

import "fmt"

func main() {
	lista := []int{1, 2, 3, 4, 5}
	var num int
	fmt.Print("escreva um numero:")
	fmt.Scan(&num)
	encontrado := false
	for _, n := range lista {
		if n == num {
			encontrado = true
		}

	}
	if encontrado == false {
		lista = append(lista, num)

	}
	fmt.Println(lista)
}
