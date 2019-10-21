package piscine

import "fmt"

func IterativeFactorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * IterativeFactorial(n-1)

}
