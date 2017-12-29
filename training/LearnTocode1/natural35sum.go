package natural35sum

import "fmt"

func main() {
	var answer int

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			answer += i
			fmt.Println(answer)
		}
	}
}
