package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ans := make([]int, max-min)
	for i := range ans {
		ans[i] = i + min
	}
	return ans

}
