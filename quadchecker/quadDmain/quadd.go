package main

import (
	"fmt"
	"os"
	"strconv"

	"quadchecker"
)

func main() {
	args := os.Args[0:]
	arg1, _ := strconv.Atoi(args[1])
	arg2, _ := strconv.Atoi(args[2])
	shape := quadchecker.QuadD(arg1, arg2)

	for _, s := range shape {
		fmt.Println(s)
	}
}
