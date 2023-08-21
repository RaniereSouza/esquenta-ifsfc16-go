package main

import (
	"fmt"
	"os"
)

func main() {

	args, name := os.Args[1:], "World"

	if len(args) > 0 {
		name = args[0]
	}

	fmt.Println("Hello, " + name + "!")
}
