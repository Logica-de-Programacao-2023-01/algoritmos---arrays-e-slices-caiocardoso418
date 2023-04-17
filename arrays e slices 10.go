package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var max, min int

	min = slice[0]
	max = slice[0]

	for _, i := range slice {

		if i < min {
			min = i
		}
		if i > max {
			max = i
		}

	}
	fmt.Println("O valor máximo é:", max)
	fmt.Println("O valor mínimo é:", min)
}
