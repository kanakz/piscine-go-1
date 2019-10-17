package main

import "github.com/01-edu/z01"

func main() {
	variableA := 'A'
	var variableB rune = 'B'

	z01.PrintRune(variableA)
	z01.PrintRune(variableB)

	z01.PrintRune('\n')

	z01.PrintRune(65)
	z01.PrintRune(66)
	z01.PrintRune(65+25)
}