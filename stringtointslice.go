package piscine

func StringToIntSlice(str string) []int {
	var newarr []int
	for _, ch := range str {
		newarr = append(newarr, int(ch))
	}
	return newarr
}
