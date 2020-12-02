package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5, 7, 6, 4, 2, 1, 4, 6}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
