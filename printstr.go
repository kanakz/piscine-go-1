package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		z01.PrintRune(rune(str1[i]))
	}
	z01.PrintRune('\n')
}
