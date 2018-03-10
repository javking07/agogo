package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}
func gen(q chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 30; i++ {
			c <- i
		}
		close(c)
		q <- 1
		close(q)
	}()

	return c
}
func receive(cc <-chan int, qq chan int) {
	for {
		select {
		case v1 := <-cc:
			fmt.Println("value is", v1)
		case <-qq:
			return
		}
	}
}
