package piscine

func BasicJoin(strs []string) string {
	str := ""
	for _, s := range strs {
		str = str + s
	}
	return str
}
