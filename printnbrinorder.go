package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var arr []rune
	arr = Nbr(n, arr)
	len := 0
	for i := range arr {
		if i == i {
			len++
		}
	}
	SortIntegerTable(arr, len)
	for i := 0; i < len; i++ {
		z01.PrintRune(arr[i])
	}
}

func SortIntegerTable(table []rune, len int) {
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if table[j] > table[i] {
				table[i] = table[i] + table[j]
				table[j] = table[i] - table[j]
				table[i] = table[i] - table[j]
			}
		}
	}
	return
}

func Nbr(n int, arr []rune) []rune {
	x := '0'
	if n == 0 {
		arr = append(arr, x)
		return arr
	} else if n > 0 {
		for i := 1; i <= n%10; i++ {
			x++
		}
		if n/10 != 0 {
			arr = append(arr, x)
			arr = Nbr(n/10, arr)
		} else {
			arr = append(arr, x)
		}
		return arr
	}
	return arr
}
