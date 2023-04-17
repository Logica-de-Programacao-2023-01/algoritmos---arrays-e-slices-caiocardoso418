package main

import "fmt"

func main() {
	slice := []string{"robson", "jorge", "cliton", "kleber", "caio", "ana", "clea", "maria"}
	fmt.Println(slice)

	var nome string
	fmt.Print("Informe um nome que queira retirar: ")
	fmt.Scan(&nome)

	novoSlice := []string{}
	for _, s := range slice {
		if s != nome {
			novoSlice = append(novoSlice, s)
		}
	}

	fmt.Println(novoSlice)
}
