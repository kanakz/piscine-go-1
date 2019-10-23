package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for index, argument := range arguments {
		if index != 0 {
			for _, c := range argument {
				z01.PrintRune(c)
			}
		}
		z01.PrintRune('\n')
	}
}
