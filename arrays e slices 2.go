package main

import "fmt"

func main() {
	slices := []int{1, 2, 3, 4, 5}
	slices = append(slices[:2], slices[3:]...)

	fmt.Println(slices)
}
