package main

import "fmt"

func main() {
	//testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//testMap := map[string]int{"one": 1, "two": 2, "three": 3}
	type person struct {
		fname   string
		lname   string
		favFood []string
	}
	// for i, v := range testMap {
	// 	fmt.Println(i, v)
	// }

	p1 := person{
		"javin",
		"forrester",
		[]string{"pancakes", "bacon", "chicken"},
	}

	fmt.Println(p1.fname, p1.lname)
	for _, v := range p1.favFood {
		fmt.Println(v)
	}
}
