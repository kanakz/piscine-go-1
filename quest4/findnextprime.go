package piscine

func Prime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for {
		if Prime(nb) {
			return nb
		}
		nb++
	}
}

// func Prime(nb int) bool {
// 	k := 0
// 	for i := 1; i <= nb; i++ {
// 		if nb%i == 0 {
// 			k++
// 		}
// 	}
// 	if k == 2 {
// 		return true
// 	}
// 	return false
// }

// func FindNextPrime(nb int) int {
// 	if Prime(nb) {
// 		return nb
// 	} else {
// 		for i := nb; i <= 1000000000; i++ {
// 			nb = nb + 1
// 			if Prime(nb) {
// 				return nb
// 			}
// 		}
// 	}
// 	return 0
// }
