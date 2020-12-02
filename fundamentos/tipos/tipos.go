package main

import (
	"fmt"
	"reflect"
)

func main() {
	//numeros inteirois
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(2))

	//sem sinal... uint8 uint16 uint16 uint32 uint64
	var b byte = 123
	fmt.Println(b)

	noma := '!'
	fmt.Println(noma)
}
