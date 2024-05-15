package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsValidAnswer() string {
	var ans string
	flag := false
	for !flag {
		fmt.Scanln(&ans)
		ans = strings.ToLower(ans)
		if ans == "yes" || ans == "no" {
			return ans
		}
		fmt.Print("Please enter 'yes' or 'no'.")
	}
	return ans
}

func IsValidInt() int {
	var num string
	flag := false
	var mynum int
	for !flag {
		fmt.Scanln(&num)
		var err error
		mynum, err = strconv.Atoi(num)
		if err != nil || mynum < 1 || mynum > 9 {
			fmt.Print("Please enter a valid number (1-9)")
			continue
		}
		break
	}
	return mynum
}

func printBoard(board [][]string) {
	fmt.Println(".-=-=--=-=-==-=-=-=-=-=-.")
	for i, row := range board {
		if i == 3 || i == 6 {
			fmt.Println("|-------+-------+-------|")
		}
		fmt.Print("|", " ")
		for j, c := range row {
			if j == 3 || j == 6 {
				fmt.Print("|", " ")
				fmt.Printf("%s", c)
				fmt.Print(" ")
			} else {
				fmt.Printf("%s", c)
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}

// Function to check the validity of placing 'c' in board
func isValid(board [][]string, row, col int, c rune) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == string(c) || board[row][i] == string(c) ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == string(c) {
			return false
		}
	}
	return true
}

// Function to solve Sudoku using backtracking
func solve(board [][]string) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == "." {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, c) {
						board[i][j] = string(c)
						if solve(board) {
							return true
						}
						board[i][j] = "."
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	flag := true
	board := make([][]string, 9)
	for i := range board {
		board[i] = make([]string, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = "."
		}
	}
	args := os.Args[1:] // Read command-line arguments
	// Check for exactly 9 arguments
	if len(args) != 9 {
		fmt.Println("Error: Please provide exactly 9 rows of the Sudoku puzzle as arguments.")
		return
	}
	for i := range board {
		for j, c := range args[i] {
			if c != '.' && (c < '1' || c > '9') {
				fmt.Printf("Invalid character '%c' in row %d.\n", c, i+1)
				return
			}
			board[i][j] = string(c)
		}
	}
	for flag {
		printBoard(board)
		fmt.Print("Do you want to enter numbers? (yes/no)")
		myans := IsValidAnswer()
		if myans == "no" {
			if solve(board) {
				fmt.Println("Sudoku solved successfully:")
				printBoard(board)
				return
			} else {
				fmt.Println("No solution exists for the provided Sudoku puzzle")
				return
			}
		}
		fmt.Print("Enter the number.")
		mynewnum := IsValidInt()
		fmt.Print("Enter the row.")
		myrow := IsValidInt()
		myrow -= 1
		fmt.Print("Enter the col.")
		mycol := IsValidInt()
		mycol -= 1
		myres := strconv.Itoa(mynewnum)
		board[myrow][mycol] = myres
	}
}
