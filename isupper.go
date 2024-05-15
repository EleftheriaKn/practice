package piscine

func isChar(k byte) bool {
	return k >= 'A' && k <= 'Z'
}

func IsUpper(s string) bool {
	new := []rune(s)
	flag := true
	for i := 0; i < len(new); i++ {
		bytenew := byte(new[i])
		if isChar(bytenew) {
			continue
		} else {
			flag = false
			break
		}
	}
	return flag
}
