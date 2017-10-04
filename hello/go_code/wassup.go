package main

import (
	"fmt"
)

var greet = "Hey, Wassup, Hello!,"
var name string

func main() {
	fmt.Print("What's your name?")
	fmt.Scanf("%s", &name)
	fmt.Println(greet, name)
}
