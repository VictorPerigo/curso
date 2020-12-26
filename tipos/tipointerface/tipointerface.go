package main

import "fmt"

type cursor struct {
	nome string
}

func main() {
	var coisa interface{}

	coisa = 3
	fmt.Println(coisa)

}
