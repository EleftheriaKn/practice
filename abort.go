package piscine

func Abort(a, b, c, d, e int) int {
	newarr := []int{a, b, c, d, e}
	for j := 0; j < 4; j++ {
		if newarr[j] > newarr[j+1] {
			newarr[j], newarr[j+1] = newarr[j+1], newarr[j]
		}
	}
	return newarr[2]
}
