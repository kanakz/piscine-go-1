package piscine

func AppendRange(min, max int) []int {
	var ans []int
	// ans := make([]int, size)
	for i := min; i < max; i++ {
		// ans[i] = ans[i] + 1
		ans = append(ans, i)
	}
	return ans
}
