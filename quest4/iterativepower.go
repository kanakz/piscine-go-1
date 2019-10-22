package piscine

func IterativePower(nb int, power int) int {
	if nb < -20 || nb > 20 {
		return 0
	}
	ans := nb
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	if power > 1 {
		for i := 2; i <= power; i++ {
			ans = ans * nb
		}
		return ans
	}
	return 0
}
