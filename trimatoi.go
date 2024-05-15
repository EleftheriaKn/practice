package piscine

func TrimAtoi(s string) int {
	integ := 0
	flag := false
	for _, i := range s {
		if i == '-' && integ == 0 {
			flag = true
		}
		if i >= '0' && i <= '9' {
			integ *= 10
			integ += int(i - '0')
		}
	}
	if flag {
		integ *= -1
	}
	return integ
}
