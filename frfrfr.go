package piscine

import (
	"github.com/01-edu/z01"
)

// PrintCombN prints all combinations of digits of length n
func PrintCombN(n int) {
	if n <= 0 || n > 9 {
		return
	}
	// Create a slice to store the current combination of digits
	current := make([]rune, n)
	printCombNRecursive(current, 0, n)
	z01.PrintRune('\n')
}

// printCombNRecursive generates combinations using recursion
func printCombNRecursive(current []rune, index, n int) {
	if index == n {
		// Print the current combination
		for i := 0; i < n; i++ {
			z01.PrintRune(current[i])
			if i < n-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Start from '0' up to '9' for the current position in the combination
	start := '0'
	if index == 0 {
		start = '1' // The first digit should start from '1' to avoid leading zeros
	}

	for i := start; i <= '9'; i++ {
		current[index] = i
		printCombNRecursive(current, index+1, n)
	}
}
