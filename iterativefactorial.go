package piscine

func IterativeFactorial(nb int) int {
	if nb < -20 || nb > 20 {
		return 0
	}
	ans := 1
	if nb <= 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			ans = ans * i
		}
	}
	return ans
}
