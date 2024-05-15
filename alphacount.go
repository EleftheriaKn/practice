package piscine

func isLatin(k byte) bool {
	return (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z')
}

func AlphaCount(s string) int {
	new := []rune(s)
	c := 0
	for i := 0; i < len(s); i++ {
		bytec := byte(new[i])
		if isLatin(bytec) {
			c += 1
		}
	}
	if c > 0 {
		return c
	}
	return 0
}
