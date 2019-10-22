package piscine

func Count(s []rune) int {
	k := 0
	for index := range s {
		index = index
		k++
	}
	return k
}

func Index(s string, toFind string) int {
	s1 := []rune(s)
	s2 := []rune(toFind)
	match := 1

	for index, letter := range s1 {
		if letter == s2[0] && Count(s1) >= Count(s2)+index-1 {
			for i := 1; i < Count(s2); i++ {
				if s2[i] == s1[index+i] {
					match++
				}
			}
			if match == Count(s2) {
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
