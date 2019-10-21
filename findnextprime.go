package piscine

func Prime(nb int) bool {
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

func FindNextPrime(nb int) int {
	if Prime(nb) {
		return nb
	} else {
		for i := nb; i <= 1e11; i++ {
			nb = nb + 1
			if Prime(nb) {
				return nb
			}
		}
	}
	return 0
}
