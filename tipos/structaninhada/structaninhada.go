package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (pedido pedido) valorTotal() (total float64) {
	total = 0.00
	for _, item := range pedido.itens {
		total += item.preco * float64(item.quantidade)
	}
	return
}

func main() {
	jonas := pedido{
		userID: 1,
		itens: []item{
			item{
				produtoID:  1,
				quantidade: 2,
				preco:      12.10,
			},
			item{
				produtoID:  2,
				quantidade: 1,
				preco:      23.49,
			},
			item{
				produtoID:  11,
				quantidade: 100,
				preco:      3.13,
			},
		},
	}

	fmt.Printf("valor Total do pedido é %.2f", jonas.valorTotal())
}
