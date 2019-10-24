package piscine

func AlphaCount(str string) int {
	count := 0
	runes := []rune(str)
	for _, v := range runes {
		if (v >= 'A' && v < 91) || (v > 96 && v <= 'z') {
			count++
		}
	}
	return count
}
