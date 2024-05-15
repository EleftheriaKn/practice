package piscine

func ToLower(s string) string {
	upper := ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			upper += string(i + 32)
		} else {
			upper += string(i)
		}
	}
	return upper
}
