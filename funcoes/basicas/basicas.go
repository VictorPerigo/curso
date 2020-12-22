package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(pedro string, pessanha string) {
	fmt.Println(pedro, pessanha)
}

func f3(p1, p2 string) string {
	return fmt.Sprintf("F3: %v %v", p1, p2)
}

func f4() (string, string) {
	return "cu", "cu2"
}

func main() {
	f1()
	f2("cu", "jonas")
	fmt.Println(f3("cu", "jonas"))
	fmt.Println(f4())
}
