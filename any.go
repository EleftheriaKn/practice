package piscine

func Any(f func(string) bool, a []string) bool {
	flag := false
	for _, ch := range a {
		if f(ch) {
			flag = true
		}
	}
	return flag
}
