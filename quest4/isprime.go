package piscine

func IsPrime(nb int) bool {
	k := 0
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			k++
		}
	}
	if k == 2 {
		return true
	}
	return false
}
