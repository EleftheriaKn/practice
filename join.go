package piscine

func Join(strs []string, sep string) string {
	newarr := strs[0]
	for i := 1; i < len(strs); i++ {
		newarr += sep
		newarr += strs[i]
	}
	return newarr
}
