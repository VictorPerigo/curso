package main

import "fmt"

type nota float64

func (nota nota) entre(inicio, fim float64) bool {
	return float64(nota) >= inicio && float64(nota) <= fim
}

func notaParaConceito(nota nota) string {
	switch {
	case nota.entre(9.0, 10.0):
		return "A"
	case nota.entre(7.0, 8.99):
		return "B"
	case nota.entre(5.0, 7.99):
		return "C"
	case nota.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
