package main

import "fmt"

func main() {
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	novo := make([]int, 0, len(lista))

	for _, tamanho := range lista {
		if tamanho > 5 {
			novo = append(novo, tamanho)
		}
	}
	fmt.Println(novo)

}
