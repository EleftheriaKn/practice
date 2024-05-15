package piscine

func LastRune(s string) rune {
	new := []rune(s)
	return new[len(s)-1]
}
