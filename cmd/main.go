package main

import (
	"fmt"
	"os"
)

// using lib as cli tool
func main() {
	args := os.Args

	for i := 1; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
