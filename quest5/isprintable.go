package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, letter := range s {
		if letter >= ' ' && letter <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
