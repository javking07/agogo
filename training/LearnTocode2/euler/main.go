package main

import (
	"fmt"
)

func nextFibNum(num1 int, num2 int) int {
	// returns sum of two ints, for next fib

	return num1 + num2
}

func iseven(number int) bool {
	// decide if int is even or not

	if number%2 == 0 {

		return true
	}

	return false
}

func main() {
	fibSeries := make([]int, 2)
	fibSeries[0] = 1
	fibSeries[1] = 2
	fibSum := fibSeries[1]

	for i := 0; i >= 0; i++ {
		fibSeries = append(fibSeries, nextFibNum(fibSeries[i], fibSeries[i+1]))

		if (nextFibNum(fibSeries[i], fibSeries[i+1])) > 4000000 {
			break
		} else if iseven(nextFibNum(fibSeries[i], fibSeries[i+1])) {
			fibSum += fibSeries[len(fibSeries)-1]
		}
	}
	fmt.Println("First 10 fib numbers: ", fibSeries[:10])
	fmt.Printf("%v%v\n", "Sum of even fib numbers up to 4 Million: ", fibSum)
	fmt.Println("Last fib number in this series is: ", fibSeries[len(fibSeries)-2])
}

/* By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.

*/
