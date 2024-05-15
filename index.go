package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		flag := true
		if s[i] == toFind[0] {
			j := 1
			for j < len(toFind) {
				if s[i+j] != toFind[j] {
					flag = false
				}
				j += 1
			}
			if flag {
				return i
			}
		}
	}
	return -1
}
