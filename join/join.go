package piscine

func Join(strs []string, sep string) string {
	news := strs[0]
	for _, i := range strs[1:] {
		news += sep + string(i)
	}
	return news
}
