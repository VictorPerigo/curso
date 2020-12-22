package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for u := 0; u <= 10; u++ {
			fmt.Printf("%v X %v = %v \n", i, u, i*u)
		}
		fmt.Printf("\n")
	}
}
