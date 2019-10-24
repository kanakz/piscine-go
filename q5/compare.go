package piscine

func Len(s string) int {
	len := 0
	for i := range s {
		if i == i {
			len++
		}
	}
	return len
}

func Compare(a, b string) int {
	// lenb := Len(b)
	// lena := Len(a)
	// c := 0
	// for _, j := range b {
	// 	for i2, j2 := range a {
	// 		if j == j2 {
	// 			for k := 0; k < lenb; k++ {
	// 				if a[i2+k] == b[k] {
	// 					c++
	// 				}
	// 			}
	// 			if c == lena {
	// 				return 0
	// 			} else if c < lena && i2 > 0 {
	// 				return -1
	// 			} else {
	// 				return 1
	// 			}
	// 		}
	// 		if c <= 0 {
	// 			return -1
	// 		}
	// 	}
	// }
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
	return 0
}
