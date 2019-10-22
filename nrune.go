package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	ans := '0'
	for index, char := range s1 {
		if index == n-1 {
			ans = char
		}
	}
	return ans
}
