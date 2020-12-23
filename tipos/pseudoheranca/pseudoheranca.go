package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //campo anonimo
	turboligado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboligado = true
	fmt.Printf("A a ferrari %v está com o turbo ligado? %v", f.nome, f.turboligado)
}
