package main

import (
	"fmt"
)

func nextFibNum(num1 int, num2 int) int {
	return num1 + num2
}

func iseven(number int) bool {
	if number%2 == 0 {
		return true
	}

	return false
}

func main() {
	fibSeries := []int{1, 2}
	fibSum := fibSeries[1]
	fmt.Println(fibSeries)
	fmt.Println(fibSum)

	/* 	for i := 0; i >= 0; i++ {
	   		if fibSeries[i]+fibSeries[i+1] >= 4000000 {
	   			break
	   		} else if iseven(nextFibNum(fibSeries[i], fibSeries[i+1])) {
	   			fibSeries = append(fibSeries, nextFibNum(fibSeries[i], fibSeries[i+1]))
	   			fibSum += fibSeries[len(fibSeries)-1]
	   		}
	   	}

	   	fmt.Printf("%v\n", fibSum)
	   	fmt.Println(fibSeries[len(fibSeries)-1])
	*/
}
