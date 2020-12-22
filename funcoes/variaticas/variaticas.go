package main

import "fmt"

var media = func(notas ...float32) float32 {
	var total float32 = 0.0
	for _, nota := range notas {
		total += nota
	}
	return total / float32(len(notas))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
