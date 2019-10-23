package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for _, c := range argument[0] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
