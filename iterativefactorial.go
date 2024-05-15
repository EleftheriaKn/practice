package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := 1
		for i := 2; i <= nb; i++ {
			if (1<<63-1)/i < result {
				return 0
			}
			result *= i
		}
		return (result)
	}
}
