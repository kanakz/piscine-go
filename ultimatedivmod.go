package piscine

func ultimateDivMod(a *int, b *int) {
	c := *a
	*a = *a/(*b)

	*b = c%(*b)
}
