package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i, s := range strs {
		if i > 0 {
			str = str + sep + s
		} else {
			str = str + s
		}
	}
	return str
}
