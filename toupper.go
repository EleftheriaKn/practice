package piscine

func ToUpper(s string) string {
	lower := ""
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			lower += string(i - 32)
		} else {
			lower += string(i)
		}
	}
	return lower
}
