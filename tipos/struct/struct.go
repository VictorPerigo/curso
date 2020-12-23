package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() (retorno float64) {
	retorno = p.preco * (1 - p.desconto)
	return
}

func (p produto) flood(palavra string, valor int) (flood string) {
	for i := 1; i <= valor; i++ {
		flood += palavra
	}
	return
}

func main() {
	produto1 := produto{
		nome:     "lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	produto2 := produto{
		nome:     "caneca",
		preco:    10.00,
		desconto: 0.5,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2, produto2.precoComDesconto(), produto2.flood("jonas", 5))
}
