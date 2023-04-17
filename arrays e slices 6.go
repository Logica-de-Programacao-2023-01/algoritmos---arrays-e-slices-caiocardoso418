package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var numeros int
	fmt.Print("escreva os numero dos arrays")
	fmt.Scan(&numeros)
	tiver := false
	for _, num := range array {
		if num == numeros {
			tiver = true

		}

	}
	if tiver {
		fmt.Println("o valor esta no array", numeros)
	} else {
		fmt.Println("o valor nao esta no array", numeros)

	}
}
