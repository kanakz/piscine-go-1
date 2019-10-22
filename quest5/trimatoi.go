package piscine

func IntCheck(s string) bool {
	ok := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			ok = true
			break
		}
	}
	return ok
}

func MinusCheck(s string) bool {
	minus := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			break
		}
		if c == '-' {
			minus = true
		}
	}
	return minus
}

func TrimAtoi(s string) int {
	x := 0
	if IntCheck(s) == true {
		for _, c := range s {
			if c >= '0' && c <= '9' {
				k := 0
				for i := '1'; i <= c; i++ {
					k++
				}
				x = x*10 + k
			}
		}
	}
	if MinusCheck(s) {
		x *= -1
	}
	return x
}

// package piscine

// import "github.com/01-edu/z01"

// func TrimAtoi(s string) int {
// 	num := []rune{}
// 	len := 0
// 	positive := false
// 	negative := false
// 	for _, l := range s {
// 		if l == '-' && !positive {
// 			negative = true
// 		}
// 		if l > 47 && l < 58 {
// 			num = append(num, l)
// 			len++
// 			if !negative {
// 				positive = true
// 			}
// 		}
// 	}
// 	d := 1
// 	ans := 0
// 	for i := len - 1; i > -1; i-- {
// 		ans = ans + (int(num[i])-48)*d
// 		d = d * 10
// 	}
// 	if negative {
// 		ans = ans * (-1)
// 	}
// 	return ans
// }

// package piscine

// func TrimAtoi(s string) int {
// 	a := 0
// 	if s == "" {
// 		return a
// 	} else {
// 		slice := []rune(s)
// 		k := 1
// 		for _, letter := range slice {
// 			if letter >= '0' && letter <= '9' {
// 				for index := range slice {
// 					if slice[index] == 45 || slice[index] == 43 {
// 						slice[index] = 's'
// 					}
// 				}
// 				break
// 			}
// 			if letter == 45 {
// 				k = -1
// 				for index := range slice {
// 					if slice[index] == 45 || slice[index] == 43 {
// 						slice[index] = 's'
// 					}
// 				}
// 				break
// 			}
// 		}
// 		for _, bukva := range slice {
// 			if bukva >= '0' && bukva <= '9' {
// 				q := 0
// 				for i := '0'; i < bukva; i++ {
// 					q++
// 				}
// 				a = a*10 + q
// 			}
// 		}
// 		a = a * k
// 		return a
// 	}
// }
