package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "guy",
	}

	a, err := json.Marshal(p1)

	fmt.Println(string(a), err)

}
