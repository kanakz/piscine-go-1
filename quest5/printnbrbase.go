package piscine

import (
	"github.com/01-edu/z01"
)

func BaseSize(t string) int {
	k := 0
	for range t {
		k++
	}
	return k
}

func BaseCheck(t string) bool {
	check := true
	m := map[rune]bool{}
	for _, c := range t {
		if c == '-' || c == '+' {
			check = false
		}
		if m[c] == false {
			m[c] = true
		} else {
			check = false
		}
	}
	return check
}

func PrintNbrBase(s int, t string) {
	max_div := BaseSize(t)
	base_len := BaseSize(t)
	ans := ""
	if s < 0 {
		ans = "-"
		max_div *= -1
	}
	if base_len > 1 && BaseCheck(t) {
		for s/max_div >= base_len {
			max_div *= base_len
		}
		for max_div != 0 {
			ans = ans + string(t[s/max_div])
			s = s - s/max_div*max_div
			max_div /= base_len
		}
	} else {
		ans = "NV"
	}

	for _, c := range ans {
		z01.PrintRune(c)
	}
}

// func PrintNbrBase(s int, t string) {
// 	ans := ""
// 	ln := 0
// 	for _, c := range t {
// 		if c == c {
// 			ln++
// 		}
// 	}
// 	mx_p := ln
// 	if s < 0 {
// 		ans = "-"
// 		mx_p *= -1
// 	}
// 	if ln > 1 {

// for s/mx_p >= ln {
// 	mx_p *= ln
// }
// for mx_p != 0 {
// 	ans = ans + string(t[s/mx_p])
// 	s = s - s/mx_p*mx_p
// 	mx_p /= ln
// }
// 		x := map[rune]bool{}
// 		for _, c := range t {
// 			if c == '+' || c == '-' {
// 				ans = "NV"
// 				break
// 			}
// 			if x[c] == false {
// 				x[c] = true
// 			} else {
// 				ans = "NV"
// 				break
// 			}
// 		}
// 	} else {
// 		ans = "NV"
// 	}
// 	for _, c := range ans {
// 		z01.PrintRune(c)
// 	}
// }
