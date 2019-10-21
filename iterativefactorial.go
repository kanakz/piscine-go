package piscine

import "github.com/01-edu/z01"

func z01.IterativeFactorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * z01.IterativeFactorial(n-1)

}
