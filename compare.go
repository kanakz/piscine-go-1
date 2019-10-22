package piscine

func Compare(a, b string) int {
	k1 := 0
	k2 := 0

	for i := range a {
		i = i
		k1++
	}
	for j := range b {
		j = j
		k2++
	}
	if k1 == k2 {
		return 0
	} else if k1 > k2 {
		return 1
	} else {
		return -1
	}
}
