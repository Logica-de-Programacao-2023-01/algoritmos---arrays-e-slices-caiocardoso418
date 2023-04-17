package main

import "fmt"

func main() {
	array := [4]float64{2.5, 2.5, 2.5, 4.5}
	soma := array[0] * array[1] * array[2] * array[3]
	fmt.Println(soma)
}
