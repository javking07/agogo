package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
   write a program that
   launches 10 goroutines
   each goroutine adds 10 numbers to a channel
   pull the numbers off the channel and print them
*/

func main() {
	workers := 10
	ch1 := make(chan int)
	counter := 0

	go add(ch1, workers)

	takeoff(ch1, &counter)
	fmt.Println("chan 1 is length of", len(ch1))
	fmt.Println(counter)

}

func add(ch chan int, work int) {
	var wg sync.WaitGroup
	for i := 0; i <= work; i++ {
		go func() {
			wg.Add(1)
			for i := 0; i <= 10; i++ {
				ch <- i
			}
			wg.Done()
			fmt.Println("number of GRs is", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	close(ch)
}

func takeoff(ch chan int, count *int) {
	for j := range ch {
		*count += j
		fmt.Println("counter value is", *count)
	}

}
