package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ca := 0
	cb := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ca += 1
		} else if f(a[i], a[i+1]) <= 0 {
			cb += 1
		}
	}
	if ca == len(a)-1 || cb == len(a)-1 {
		return true
	}
	return false
}
