package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filename := os.Args[1]
	ch, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%s", ch)
	return
}
