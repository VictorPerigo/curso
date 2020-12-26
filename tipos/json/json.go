package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := product{1, "notebook", 1899.90, []string{"promoção", "eletronico"}}
	p1json, _ := json.Marshal(p1)
	fmt.Println(string(p1json))

	//json apara struct
	p2 := product{}
	jsonString := `{"id":2,"nome":"caneta","preco":8.90,"tags":["Papelaria"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0])
}
