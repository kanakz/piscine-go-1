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

func Len(d string) int {
	inc := 0
	for range d {
		inc++
	}
	return inc
}

func Reverse(s string) string {
	var reverse string
	for i := Len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func NbrBase(nbr int, baseFrom string) string {
	indx := 0
	for _, res := range baseFrom {
		if string(res) == "-" || string(res) == "+" {
			indx = 1
			break
		}
	}
	if indx == 1 || len(baseFrom) < 2 {
		return "0"
	} else if 2147483647 < nbr || -2147483648 > nbr {
		return string(nbr)
	} else {
		i := 0
		smt := ""
		for nbr >= Len(baseFrom) {
			if nbr >= Len(baseFrom) {
				smt += string(baseFrom[nbr%Len(baseFrom)])
				nbr = nbr / Len(baseFrom)
				i++
			}
		}
		smt += string(baseFrom[nbr])
		return Reverse(smt)
	}
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

// func ConvertBase(s, t, p string) string {
// 	ln := 0
// 	ln2 := 0
// 	ln3 := 0
// 	ans := ""
// 	st_t := map[rune]int{}
// 	for c := range s {
// 		ln = c + 1
// 	}
// 	for i, c := range t {
// 		st_t[c] = i
// 		ln2 = i + 1
// 	}
// 	for c := range p {
// 		ln3 = c + 1
// 	}
// 	pw := 1
// 	cnt := 0
// 	for i := ln - 1; i >= 0; i-- {
// 		cnt = cnt + st_t[[]rune(s)[i]]*pw
// 		pw *= ln2
// 	}
// 	for cnt != 0 {
// 		ans = string(p[cnt%ln3]) + ans
// 		cnt /= ln3
// 	}
// 	return ans
// }

// func ConvertBase(nbr, baseFrom, baseTo string) string {
// 	n := AtoiBase(nbr, baseFrom)
// 	return PrintNbrBase(n, baseTo)
// }

// func AtoiBase(s string, base string) int {

// 	len_s := 0
// 	for i := range s {
// 		len_s = i + 1
// 	}

// 	len_base := 0
// 	for i := range base {
// 		len_base = i + 1
// 	}

// 	if len_base < 2 {
// 		return 0
// 	}

// 	for i := 0; i < len_base-1; i++ {
// 		for j := i + 1; j < len_base; j++ {
// 			if base[i] == base[j] {
// 				return 0
// 			}
// 		}
// 	}

// 	count := 0
// 	var nums [40]int
// 	for i := range s {
// 		for j := range base {
// 			if s[i] == base[j] {
// 				nums[count] = j
// 				count++
// 			}
// 		}
// 	}

// 	res := 0
// 	for i := 0; i < len_s; i++ {
// 		res += nums[i] * pow(len_base, len_s-1-i)
// 	}
// 	return res

// }

// func pow(a int, b int) int {
// 	if b < 0 {
// 		return 0
// 	}
// 	ans := 1
// 	for b != 0 {
// 		ans *= a
// 		b--
// 	}
// 	return ans
// }

// func PrintNbrBase(nbr int, base string) string {
// 	length := 0
// 	for i := range base {
// 		length = i + 1
// 	}

// 	if length < 2 {
// 		return ""
// 	}
// 	for i := 0; i < length-1; i++ {
// 		for j := i + 1; j < length; j++ {
// 			if base[i] == base[j] {
// 				return ""
// 			}
// 		}
// 	}

// 	var nums [70]int
// 	var runes [70]rune
// 	count := 0
// 	length_num := 0
// 	temp := nbr
// 	for temp != 0 {
// 		length_num++
// 		temp = temp / length
// 	}

// 	for nbr != 0 {
// 		rem := nbr % length
// 		nums[length_num-1-count] = rem
// 		count++
// 		nbr = nbr / length
// 	}

// 	base_runes := []rune(base)
// 	c := 0
// 	for i := 0; i < length_num; i++ {
// 		for j := range base_runes {
// 			if j == nums[i] {
// 				runes[c] = base_runes[j]
// 				c++
// 			}
// 		}
// 	}

// 	return string(runes[:length_num])
// }
