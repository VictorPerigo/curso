package main

import "fmt"

func main() {
	var cu [3]int
	s1 := cu[:]
	s2 := cu[:]

	fmt.Println(s1, s2, cu)

	s1[0] = 21

	fmt.Println(s1, s2, cu)
}
