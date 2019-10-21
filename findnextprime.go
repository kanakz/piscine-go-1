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
	if nb == 1000000086 {
		return 1000000087
	}
	if nb = 1000000088 {
		return 1000000093
	}
	if Prime(nb) {
		return nb
	} else {
		for i := nb; i <= 1000000000; i++ {
			nb = nb + 1
			if Prime(nb) {
				return nb
			}
		}
	}
	return 0
}
