package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	flag := false

	if args[0] == "--upper" {
		flag = true
		args = args[1:]
	}

	for _, ch := range args {
		for _, c := range ch {
			if c > ('9' + '9' + '8') {
				z01.PrintRune(' ')
			} else if flag {
				z01.PrintRune('A' + (c - '1'))
			} else {
				z01.PrintRune('a' + (c - '1'))
			}
		}
	}
}
