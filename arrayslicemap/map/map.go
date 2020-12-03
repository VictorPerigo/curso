package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678] = "maria"
	aprovados[21343144] = "pedro"
	aprovados[23234523] = "Anna"
	for cpf, nome := range aprovados {
		fmt.Printf("%v (CPF: %v)\n", nome, cpf)
	}
}
