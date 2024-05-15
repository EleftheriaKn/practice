package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a
	tempB := *b
	*a = tempA / tempB
	*b = tempA % tempB
}
