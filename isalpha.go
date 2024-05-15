package piscine

func IsAlpha(s string) bool {
	flag := true
	for _, i := range s {
		if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') && (i < '0' || i > '9') {
			flag = false
		}
	}
	return flag
}
