package piscine

import "github.com/01-edu/z01"

func NbrSet(x int) {
	a := '0'
	if x == 0 {
		z01.PrintRune(a)
		return
	}
	if x < 0 {
		for i := -1; i >= x%10; i-- {
			a++
		}
	}
	if x > 0 {
		for i := 1; i <= x%10; i++ {
			a++
		}
	}
	
	if x/10 != 0 {
		NbrSet(x/10)
	}
	z01.PrintRune(a)
	return
}
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	NbrSet(n)
}
