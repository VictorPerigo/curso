package main

import "fmt"

func main() {
	// a1 := [3]int{1, 2, 3} // array
	// s1 := []int{1, 2, 3}  // slice
	// fmt.Println(a1, s1)

	// a2 := [5]int{1, 2, 3, 4, 5}

	// //slice não é um array
	// s2 := a2[:]

	// fmt.Println(a2, s2)

	//slice make

	s := make([]int, 0)
	// s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}
