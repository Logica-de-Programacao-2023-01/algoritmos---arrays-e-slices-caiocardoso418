package main

import "fmt"

func main() {
	array := [6]float64{1.2, 3.4, 5.6, 7.8, 9.1, 2.3}

	var num float64
	fmt.Print("Informe um número: ")
	fmt.Scan(&num)

	for i := 0; i < len(array); i++ {
		array[i] += num
	}

	fmt.Println(array)
}
