package piscine

func Map(f func(int) bool, a []int) []bool {
	newarr := make([]bool, len(a))
	for i, ch := range a {
		newarr[i] = f(ch)
	}
	return newarr
}
