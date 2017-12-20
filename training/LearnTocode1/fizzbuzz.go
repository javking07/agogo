package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if (i%3 == 0 && i%5 == 0) { // if divisible by 3 and 5, print FizzBuzz
      fmt.Println(i," :FizzBuzz")
    } else if (i%3 == 0) { // if divisible by 3, print Fizz
      fmt.Println(i," :Fizz")
    } else if (i%5 == 0) { // if divisible by 3, print Buzz
      fmt.Println(i," :Buzz")
    }
  }
}
