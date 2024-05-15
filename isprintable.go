package piscine

func IsPrintable(s string) bool {
	flag := true
	for _, i := range s {
		if i < 32 || i > 126 {
			flag = false
		}
	}
	return flag
}
