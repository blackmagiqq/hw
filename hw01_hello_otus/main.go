package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	stringForReverse := "Hello, OTUS!"
	fmt.Println(reverse.String(stringForReverse))
}
