package main

import "fmt"

/* Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let’s us know whether or not the argument was even. For example:
half(1) returns (0, false)
half(2) returns (1, true) */

func main() {

	fmt.Println(halfEven(2))

	halfEven2 := func(number int) (int, bool) {
		return (number / 2), (number%2 == 0)
	}

	fmt.Println(halfEven2(2))
}

func halfEven(number int) (halved int, isEven bool) {

	return (number / 2), number%2 == 0
}
