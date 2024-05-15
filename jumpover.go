package piscine

func JumpOver(str string) string {
	newarr := ""
	for i := 2; i < len(str); {
		if str[i] != ' ' {
			newarr += string(str[i])
		}
		i += 3
	}
	newarr += "\n"
	return newarr
}
