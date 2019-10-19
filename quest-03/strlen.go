package piscine

func StrLen(str string) int {
	str1 := []rune(str)
	ln := 0
	for _, i := range str1 {
		if i == i {
			ln++
		}
	}
	return ln
}
