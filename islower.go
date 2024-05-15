package piscine

func isCharl(k byte) bool {
	return k >= 'a' && k <= 'z'
}

func IsLower(s string) bool {
	new := []rune(s)
	flag := true
	for i := 0; i < len(new); i++ {
		bytenew := byte(new[i])
		if isCharl(bytenew) {
			continue
		} else {
			flag = false
			break
		}
	}
	return flag
}
