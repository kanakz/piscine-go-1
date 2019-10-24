package piscine

// func Len(str string) int {
// 	k := 0
// 	for range str {
// 		k++
// 	}
// 	return k
// }

// func SplitWhiteSpaces(str string) []string {
// 	strs_el := ""
// 	strs := []string{}
// 	for index, letter := range str {
// 		if index == Len(str)-1 && string(letter) != " " && string(letter) != "\t" && string(letter) != "\n" {
// 			strs_el += string(letter)
// 			strs = append(strs, strs_el)
// 		} else if string(letter) != " " && string(letter) != "\t" && string(letter) != "\n" {
// 			strs_el += string(letter)
// 		} else {
// 			if Len(strs_el) >= 1 {
// 				strs = append(strs, strs_el)
// 			}
// 			strs_el = ""
// 		}
// 	}
// 	return strs
// }

func SplitWhiteSpaces(str string) []string {
	k := 0
	no_white_spaces := false
	for index := range str {
		if no_white_spaces && index != 0 && (str[index-1] == '\n' || str[index-1] == '\t' || str[index-1] == ' ') && str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			k++
		}
		if str[index] != '\n' && str[index] != '\t' && str[index] != ' ' {
			no_white_spaces = true
		}
	}
	k++

	x := 0
	ans := make([]string, k)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		ans[x] = ans[x] + string(c)
		ok = false
	}
	return ans
}

// func SplitWhiteSpaces(str string) []string {
// 	ss := ""
// 	len := 0
// 	prev := false
// 	for _, s := range str {
// 		if (s == ' ' || s == '\n' || s == '\t') && !prev {
// 			prev = true
// 			len++
// 		} else {
// 			prev = false
// 		}
// 	}
// 	len++

// 	arr := make([]string, len)

// 	count := 0
// 	for _, s := range str {

// 		if s == ' ' || s == '\n' || s == '\t' {
// 			length := 0
// 			for i := range ss {
// 				length = i + 1
// 			}
// 			if length == 0 {
// 				continue
// 			}
// 			arr[count] = ss
// 			count++
// 			ss = ""
// 			continue
// 		}
// 		ss += string(s)
// 	}
// 	length := 0
// 	for i := range ss {
// 		length = i + 1
// 	}
// 	if length != 0 {
// 		arr[count] = ss
// 	}
// 	return arr
// }
