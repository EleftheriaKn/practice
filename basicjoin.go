package piscine

func BasicJoin(elems []string) string {
	news := ""
	for _, i := range elems {
		news += string(i)
	}
	return news
}
