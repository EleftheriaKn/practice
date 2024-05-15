package piscine

import "fmt"

func PrintNbr(n int) {
	if n < 0 {
		if n == -9223372036854775808 {
			fmt.Print('-')
			fmt.Print('9')
			n = 223372036854775808
		} else {
			fmt.Print('-')
			n *= -1
		}
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	fmt.Print(rune(n%10 + '0'))
}
