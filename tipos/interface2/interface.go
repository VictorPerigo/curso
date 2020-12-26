package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboligado     bool
	velocidadeAtual int
}

func (ferrari *ferrari) ligarTurbo() {
	ferrari.turboligado = true
}

func main() {
	carro1 := ferrari{"f40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"f40", false, 0}
	fmt.Println(carro2)
}
