package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	str1 := []rune(str)
	for a, word := range str1 {
		z01.PrintRune(rune(word))
		a++
	}
}
