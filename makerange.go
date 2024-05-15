package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	c := 0
	s := make([]int, max-min)
	for i := min; i < max; i++ {
		s[c] = i
		c += 1
	}
	return s
}
