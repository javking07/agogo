package main

import "fmt"

func main() {

	bigsmall()

}

func bigsmall() {
	var big int
	var small int
	var answer int

	fmt.Println("Enter a big number")
	fmt.Scanf("%v", &big)
	fmt.Println("Enter a small number")
	fmt.Scanf("%v", &small)

	answer = big % small

	fmt.Println(answer)

}
