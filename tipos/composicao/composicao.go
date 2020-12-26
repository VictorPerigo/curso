package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct {
}

func (bmw7 bmw7) ligarTurbo() {
	fmt.Println("turbo ligado")
}

func (bmw7 bmw7) fazerBaliza() {
	fmt.Println("baliza ligado")
}

func main() {
	// var b esportivoLuxuoso = bmw7{}
	b := bmw7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
