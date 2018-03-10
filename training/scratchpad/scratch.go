package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(string(os.Getenv("APP_DB_USERNAME")))
}
