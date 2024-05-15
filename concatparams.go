package piscine

func ConcatParams(args []string) string {
	s := make([]string, len(args))
	news := ""
	for i, ch := range args {
		s[i] = ch
		news += s[i]
		if i < len(args)-1 {
			news += "\n"
		}
	}
	return news
}
