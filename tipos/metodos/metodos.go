package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (pessoa pessoa) getNomeCompleto() (nomeCompleto string) {
	nomeCompleto = pessoa.nome + " " + pessoa.sobrenome
	return
}

func (pessoa *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	pessoa.nome = partes[0]
	pessoa.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("maria pereira")
	fmt.Println(p1.getNomeCompleto())
}
