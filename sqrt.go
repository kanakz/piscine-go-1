package piscine

func Sqrt(nb int) int {
	r := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			r = i
		}
	}
	return r
}
