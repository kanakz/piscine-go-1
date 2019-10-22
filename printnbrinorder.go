package piscine

import (
	"github.com/01-edu/z01"
)

func IntToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return digits
}

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[j] > table[i] {
				table[i] = table[i] + table[j]
				table[j] = table[i] - table[j]
				table[i] = table[i] - table[j]
			}
		}
	}
	return table
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for _, c := range SortIntegerTable(IntToDigits(n)) {
			z01.PrintRune(rune(c) + '0')
		}
	}
}

// func IntToRune(n int) {
// 	if n <= 0 {
// 		return
// 	} else {
// 		z01.PrintRune(rune(48 + n%10))
// 		IntToRune((n - n%10) / 10)
// 	}
// }

// func RuneTransform(x int) rune {
// 	r := '0'
// 	for i := 0; i < x; i++ {
// 		r++
// 	}
// 	return r
// }

// func PrintNbrInOrder(n int) {
// 	o := 0
// 	s := []rune{}
// 	for n/10 != 0 {
// 		o = n % 10
// 		n = n / 10
// 		s = append(s, RuneTransform(o))
// 	}
// 	s = append(s, RuneTransform(n))
// 	fmt.Println(s)
// }
