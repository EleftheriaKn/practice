package piscine

func SplitWhiteSpaces(s string) []string {
	var ns []string
	var word string
	flag := false
	for _, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			if flag {
				ns = append(ns, word)
				word = ""
				flag = false
			}
		} else {
			word += string(char)
			flag = true
		}
	}
	if flag {
		ns = append(ns, word)
	}
	return ns
}
