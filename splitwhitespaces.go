package piscine

func Len(str string) int {
	k := 0
	for range str {
		k++
	}
	return k
}

func SplitWhiteSpaces(str string) []string {
	strs_el := ""
	strs := []string{}
	for index, letter := range str {
		if index == Len(str)-1 && string(letter) != " " && string(letter) != "\t" && string(letter) != "\n" {
			strs_el += string(letter)
			strs = append(strs, strs_el)
		} else if string(letter) != " " && string(letter) != "\t" && string(letter) != "\n" {
			strs_el += string(letter)
		} else {
			if Len(strs_el) >= 1 {
				strs = append(strs, strs_el)
			}
			strs_el = ""
		}
	}
	return strs
}

// func SplitWhiteSpaces(str string) []string {
// 	ln := 0
// 	ok2 := false
// 	for c := range str {

// 		if ok2 && c != 0 && (str[c-1] == '\n' || str[c-1] == '\t' || str[c-1] == ' ') && str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
// 			ln++
// 		}
// 		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
// 			ok2 = true
// 		}
// 	}
// 	ln++

// 	x := 0
// 	ans := make([]string, ln)
// 	ok := true
// 	for _, c := range str {
// 		if c == '\n' || c == '\t' || c == ' ' {
// 			if !ok {
// 				x++
// 			}
// 			ok = true
// 			continue
// 		}
// 		ans[x] = ans[x] + string(c)
// 		ok = false
// 	}
// 	return ans
// }
