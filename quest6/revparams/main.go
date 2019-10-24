package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	n := 0
	for index := range arguments {
		n = index
	}
	for i := n; i >= 1; i-- {
		for _, letter := range arguments[i] { // arguments[i] = argument, print letter by letter
			z01.PrintRune(letter) // letter type is rune
		}
		z01.PrintRune('\n')
	}
}
