package piscine

func NRune(s string, n int) rune {
	new := []rune(s)
	if n > 0 && n <= len(s) {
		return new[n-1]
	}
	return 0
}
