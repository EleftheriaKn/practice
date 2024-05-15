package piscine

func IsNumeric(s string) bool {
	flag := true
	for _, i := range s {
		if i < '0' || i > '9' {
			flag = false
		}
	}
	return flag
}
