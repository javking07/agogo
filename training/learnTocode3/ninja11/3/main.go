package main

import "fmt"

type customErr struct {
	data string
}

func (e customErr) Error() string {
	return e.data
}
func foo(e error) {
	fmt.Println(e)
}
func main() {
	x := customErr{"the error"}
	// foo(x)
	fmt.Println(x.Error())
}
