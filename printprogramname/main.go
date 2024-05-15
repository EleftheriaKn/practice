package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	lastindex := 0
	for i := len(args[0]) - 1; i >= 0; i-- {
		if args[0][i] == '/' {
			lastindex = i
			break
		}
	}
	result := args[0][lastindex+1:]
	for _, i := range result {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
