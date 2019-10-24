package piscine

func ToUpper(s string) string {
	Arune := []rune(s)
	for i := range Arune {
		if Arune[i] >= 'a' && Arune[i] <= 'z' {
			Arune[i] = Arune[i] - 32
		}
	}
	return string(Arune)
}
