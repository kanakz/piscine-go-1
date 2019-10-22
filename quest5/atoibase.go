package piscine

func AtoiBase(s string, t string) int {
	ans := 0
	ln := 0
	st := map[rune]bool{}
	for _, c := range t {
		if st[c] || c == '-' || c == '+' {
			return ans
		}
		st[c] = true
		ln++
	}
	if ln > 1 {
		for _, c := range s {
			cnt := 0
			if st[c] {
				for _, j := range t {
					if j == c {
						break
					}
					cnt++
				}
				ans = ans*ln + cnt
			}
		}
	}
	return ans
}

// package piscine

// func AtoiBase(s string, base string) int {
// 	len := 0
// 	ans := 0
// 	lens := 0
// 	for i, l := range base {
// 		len = i + 1
// 		if l == '+' || l == '-' {
// 			return 0
// 		}
// 	}
// 	if len < 2 {
// 		return 0
// 	}
// 	for i := 0; i < len-1; i++ {
// 		for j := i + 1; j < len; j++ {
// 			if base[i] == base[j] {
// 				return 0
// 			}
// 		}
// 	}
// 	for i := range s {
// 		lens = i + 1
// 	}
// 	mul := 1
// 	for i := lens - 1; i > -1; i-- {
// 		for j, cb := range base {
// 			if rune(s[i]) == cb {
// 				ans = ans + j*mul
// 				mul = mul * len
// 			}
// 		}
// 	}
// 	return ans
// }
