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
	arguments = arguments[1 : n+1]

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arguments[i] > arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}
	for _, argument := range arguments {
		for _, letter := range argument {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

// package main

// import (
// 	"github.com/01-edu/z01"
// 	"os"
// )

// func main() {
// 	arguments := os.Args
// 	k := 0
// 	for index := range arguments {
// 		index = index
// 		k++
// 	}
// 	for q := 1; q <= k-2; q++ {
// 		for i := 1; i <= k-2; i++ {
// 			if arguments[i] > arguments[i+1] {
// 				str := arguments[i+1]
// 				arguments[i+1] = arguments[i]
// 				arguments[i] = str
// 			}
// 		}
// 	}
// 	for i := 1; i <= k-1; i++ {
// 		str := arguments[i]
// 		slice := []rune(str)
// 		for _, letter := range slice {
// 			z01.PrintRune(letter)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// package main

// import "github.com/01-edu/z01"
// import "os"

// func SortIntegerTable(table []rune, len int) []rune {
// 	for i := 0; i < len; i++ {
// 		for j := 1; j < len-i; j++ {
// 			if table[j-1] > table[j] {
// 				table[j-1], table[j] = table[j], table[j-1]
// 			}
// 		}
// 	}
// 	return table
// }

// func main() {
// 	table := os.Args
// 	len := 0
// 	var arr [1000]rune
// 	for i, m := range table {
// 		if i == 0 {
// 			continue
// 		}
// 		for _, j := range m {
// 			arr[i] = j
// 			break
// 		}
// 		len = i
// 	}
// 	newarr := SortIntegerTable(arr[1:len+1], len)
// 	for _, l := range newarr {
// 		z01.PrintRune(l)
// 		z01.PrintRune(10)
// 	}
// }
