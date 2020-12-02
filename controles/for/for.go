package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Println(time.Now().UTC().Format("02-01-2006 15 00:00"))
		time.Sleep(time.Second)
	}
}
