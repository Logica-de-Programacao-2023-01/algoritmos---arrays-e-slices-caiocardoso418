package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}

	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	array[0] = num
	array[2] = num

	fmt.Println(array)
}
