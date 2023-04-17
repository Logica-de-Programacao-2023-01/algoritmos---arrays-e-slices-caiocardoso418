package main

import "fmt"

func main() {
	lista := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	novoSlice := make([]int, 0, len(lista))

	for _, i := range lista {
		if i%2 == 0 {
			novoSlice = append(novoSlice, i)
		}
	}
	fmt.Println(novoSlice)

}
