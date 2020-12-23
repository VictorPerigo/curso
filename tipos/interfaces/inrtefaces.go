package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente

func (pessoa pessoa) toString() string {
	return pessoa.nome + " " + pessoa.sobrenome
}

func (produto produto) toString() string {
	return fmt.Sprintf("%v - R$ %.2f", produto.nome, produto.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{
		nome:      "roberto",
		sobrenome: "Silva",
	}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
