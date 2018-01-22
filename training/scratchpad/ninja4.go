package main

import (
	"fmt"
)

func main() {
	s := struct {
		name       string
		age        int
		dictionary map[string]int
	}{
		name: "test",
		age:  5,
		dictionary: map[string]int{
			"javin":   1,
			"kristie": 2,
		},
	}

	fmt.Println(s)
}
