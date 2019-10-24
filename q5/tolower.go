package piscine

func ToLower(s string) string {
	Arune := []rune(s)
	for i := range Arune {
		if Arune[i] >= 'A' && Arune[i] <= 'Z' {
			Arune[i] = Arune[i] + 32
		}
	}
	return string(Arune)
}
