package piscine

import (
	"fmt"
	"strings"
)

func isValidBase(base string) bool {
	// Base must contain at least 2 characters
	if len(base) < 2 {
		return false
	}

	// Base should not contain + or - characters
	if strings.ContainsAny(base, "+-") {
		return false
	}

	// Each character of a base must be unique
	seen := make(map[rune]bool)
	for _, char := range base {
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func convertToBase(num int, base string) string {
	if num == 0 {
		return "0"
	}

	baseLen := len(base)
	result := ""
	for num > 0 {
		remainder := num % baseLen
		result = string(base[remainder]) + result
		num /= baseLen
	}

	return result
}

func PrintIntInBase(num int, base string) {
	// Check validity of the base
	if !isValidBase(base) {
		fmt.Println("NV (Not Valid)")
		return
	}

	// Handle negative numbers
	negative := num < 0
	if negative {
		num = -num
	}

	// Convert num to the specified base
	result := convertToBase(num, base)

	// Print the result, considering the sign
	if negative {
		fmt.Printf("-%s\n", result)
	} else {
		fmt.Println(result)
	}
}

func PrintMemory(arr [10]byte) {
	for i := 0; i < len(arr); i++ {
		PrintIntInBase(16, string(arr[i]))
	}
}
