package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	k := 0
	r := 0
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			k++
			r = i
		}
		return r
	}
	return 0
}
