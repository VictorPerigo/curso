package main

import "fmt"

func main() {
	aprovados := []string{"maria", "pedro", "rafael", "guilherme"}
	aprovadosEnem(aprovados...)
}

func aprovadosEnem(aprovados ...string) {
	fmt.Println("lista de aprovados")
	for indice, aprovado := range aprovados {
		fmt.Printf("aprovado de numero %v = %v \n", indice+1, aprovado)
	}
}
