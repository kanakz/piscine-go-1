package piscine

// func Count(s []rune) int {
// 	k := 0
// 	for index := range s {
// 		index = index
// 		k++
// 	}
// 	return k
// }

func Index(s string, toFind string) int {
	sS := []rune(s)
	sF := []rune(toFind)
	kS := 1
	kF := 1
	for index := range sF {
		index = index
		kS++
	}
	for index := range sS {
		index = index
		kF++
	}
	for index, letter := range sS {
		if letter == sF[0] && kF >= kS+index-1 {
			m := 1
			for i := 1; i < kS; i++ {
				if sF[i] == sS[index+i] {
					m++
				}
			}
			if m == kS {
				return index
			}
		}
	}
	return -1
}

// func Index(s string, toFind string) int {
// 	if toFind == "" {
// 		return 0
// 	}
// 	r := []rune(s)
// 	len := 0
// 	lenf := 0
// 	for i := range r {
// 		len = i
// 	}

// 	for i := range toFind {
// 		lenf = i
// 	}

// 	for i := 0; i < len-lenf; i++ {
// 		if s[i:i+lenf+1] == toFind {
// 			return i
// 		}
// 	}
// 	return -1
// }

// package quest5

// func Index(s, t string) int {
// 	ln := 0
// 	ln2 := 0
// 	for _, c := range s {
// 		if c == c {
// 			ln++
// 		}
// 	}

// 	for _, c := range t {
// 		if c == c {
// 			ln2++
// 		}
// 	}
// 	for i := 0; i < ln; i++ {
// 		if ln2 != 0 && s[i] == t[0] {
// 			ok := true
// 			cur_ch := 0
// 			for j := 0; j < ln2; j++ {
// 				if i+cur_ch >= ln || t[j] != s[i+cur_ch] {
// 					ok = false
// 					break
// 				}
// 				cur_ch++
// 			}
// 			if ok == true {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }
