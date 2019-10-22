package piscine

func AlphaCount(str string) int {
	k := 0
	for _, c := range str {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			k++
		}
	}
	return k
}
