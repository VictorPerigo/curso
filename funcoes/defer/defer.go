package main

import "fmt"

func obter(numero int) int {
	defer fmt.Println("fim!")

	if numero > 5000 {
		fmt.Println("valor alto...")
		return 5000
	}
	fmt.Println("valor baixo...")
	return numero

}

func main() {
	obter(5666666666666)
	obter(2)
}
