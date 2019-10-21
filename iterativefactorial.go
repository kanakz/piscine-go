package main

import "fmt"

func IterativeFactorial(n uint) uint{
 
    if n == 0{
        return 1
    }
    return n * IterativeFactorial(n - 1)
}
func main() {
     
    fmt.Println(IterativeFactorial(4))   // 24
    fmt.Println(IterativeFactorial(5))   // 120
    fmt.Println(IterativeFactorial(6))   // 720
}






// package piscine

// func IterativeFactorial(nb int) int {
// 	if nb < -20 || nb > 20 {
// 		return 0
// 	}
// 	if nb == 1 || nb == 0 {
// 		return 1
// 	} else if nb < 0 {
// 		return 0
// 	} else {
// 		fact := 1
// 		for i := 2; i <= nb; i++ {
// 			fact *= i
// 		}
// 		return fact
// 	}
// }