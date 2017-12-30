package main

import "fmt"

func main() {
	greatest := max(1, 2, 3, 4, 5, 6, 7, 8, 15, 10)
	fmt.Printf("%v", greatest)
}

func max(listofints ...int) int {
	var largest int
	for _, v := range listofints {
		if v > largest {
			largest = v
		}
	}
	return largest

}
