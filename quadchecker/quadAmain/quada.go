package main

import (
	"fmt"
	"os"
	"quadchecker"
	"strconv"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				fmt.Printf("o")
			} else if i == 0 && j == x-1 {
				fmt.Printf("o")
			} else if i == y-1 && j == 0 {
				fmt.Printf("o")
			} else if i == y-1 && j == x-1 {
				fmt.Printf("o")
			} else if i == 0 || i == y-1 {
				fmt.Printf("-")
			} else if j == 0 || j == x-1 {
				fmt.Printf("|")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	args := os.Args[0:]
	arg1, _ := strconv.Atoi(args[1])
	arg2, _ := strconv.Atoi(args[2])
	shape := quadchecker.QuadA(arg1, arg2)

	for _, s := range shape {
		fmt.Println(s)
	}
}
