package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	takeinput()
	inputtosum := grabinput()
	fmt.Println("second item came out as", stringsum(inputtosum))
}

func takeinput() {
	var input string
	fmt.Scanln(&input)
}
func grabinput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input: ")
	input, _ := reader.ReadString('\n')
	return input
}

func stringsum(x string) int {
	strtolist := strings.Split(x, " ")
	var iterator int
	total := 0
	for _, v := range strtolist {
		v = strings.TrimSpace(v)
		iterator, _ := strconv.Atoi(v)
		total += iterator

	}
	return total + iterator
}
