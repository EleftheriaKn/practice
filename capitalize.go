package piscine

func Capitalize(s string) string {
	news := ""
	if s[0] >= 'a' && s[0] <= 'z' {
		news += string(s[0] - 32)
	} else {
		news += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') && (news[i-1] < 'A' || news[i-1] > 'Z') && (news[i-1] < '0' || news[i-1] > '9') && (news[i-1] < 'a' || news[i-1] > 'z') {
			news += string(s[i] - 32)
		} else {
			if (s[i] >= 'A' && s[i] <= 'Z') && ((news[i-1] >= 'a' && news[i-1] <= 'z') || (news[i-1] >= 'A' && news[i-1] <= 'Z') || (s[i-1] >= '0' && s[i-1] <= '9')) {
				news += string((s[i] + 32))
			} else {
				news += string(s[i])
			}
		}
	}
	return news
}
