package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	workers := 33
	var counter int64
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			//runtime.Gosched()
			atomic.AddInt64(&counter, 1)
			fmt.Println("Goroutines:", runtime.NumGoroutine(), "Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines Ending:", runtime.NumGoroutine())
	fmt.Println("counter ending:", counter)
}
