package main

import "fmt"

func imprimirResultado(nota float32) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func main() {
	imprimirResultado(8)
}
