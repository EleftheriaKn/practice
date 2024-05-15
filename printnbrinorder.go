package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		return
	}
	digits := countDigits(n)
	printDigits(digits)
}

func countDigits(n int) [10]int {
	var digits [10]int
	for n > 0 {
		digits[n%10]++
		n /= 10
	}
	return digits
}

func printDigits(digits [10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			z01.PrintRune(rune('0' + i))
		}
	}
}
