package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	c := 0
	for c*c <= nb {
		if c*c == nb {
			return c
		}
		c += 1
	}
	return 0
}
