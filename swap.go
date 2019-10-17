package piscine

func Swap(a *int, b *int) {
	x := *a
	y := *b
	*b = x
	*a = y
}
