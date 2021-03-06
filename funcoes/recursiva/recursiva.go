package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0 || n == 1:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	result := fatorial(4)

	fmt.Println(result)
}
