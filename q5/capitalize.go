package piscine

func Capitalize(s string) string {
	Arune := []rune(s)
	len := 0
	for i := range s {
		if i == i {
			len++
		}
	}
	for i := range Arune {
		if i == 0 {
			if Arune[i] >= 'a' && Arune[i] <= 'z' {
				Arune[i] = Arune[i] - 32
			}
		} else if i+1 <= len && i > 0 {
			if Arune[i-1] >= '0' && Arune[i-1] <= '9' {
				if Arune[i] >= 'A' && Arune[i] <= 'Z' {
					Arune[i] = Arune[i] + 32
				}
			} else if (Arune[i-1] >= 'A') && (Arune[i-1] <= 'Z') || (Arune[i-1] >= 'a') && (Arune[i-1] <= 'z') {
				if Arune[i] >= 'A' && Arune[i] <= 'Z' {
					Arune[i] = Arune[i] + 32
				}
			} else {
				if Arune[i] >= 'a' && Arune[i] <= 'z' {
					Arune[i] = Arune[i] - 32
				}
			}
		}
	}
	return string(Arune)
}
