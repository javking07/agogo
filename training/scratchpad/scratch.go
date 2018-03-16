package main

import (
	"flag"
	"fmt"
)

func main() {
	var flag1 = flag.String("flag1", "flag 77777", "the first flag")
	var flag2 = flag.String("flag2", "flag 2", "the second flag")
	//flag.Parse()
	fmt.Println(*flag1)
	fmt.Println(*flag2)
}
