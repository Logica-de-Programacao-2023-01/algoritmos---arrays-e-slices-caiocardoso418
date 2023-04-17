package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)
	soma := 0

	for i := range slice {
		fmt.Printf("Informe o valor do elemento: ", i+1)
		fmt.Scanln(&slice[i])
		soma += slice[i]
	}

	fmt.Println("Slice:", slice)
	fmt.Println("Soma dos valores:", soma)
}
