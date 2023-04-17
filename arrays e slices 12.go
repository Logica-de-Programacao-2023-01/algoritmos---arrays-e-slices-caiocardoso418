package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	var slice []int
	for num := range array {
		if num%3 == 0 {
			slice = append(slice, num)

		}
	}
	fmt.Println(slice)
}
