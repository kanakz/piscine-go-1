package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x <= 0 && y <= 0 {
		z01.PrintRune('\n')
	}

	for i := 1; i <= x; i++ {
		if i == x || i == 1 {
			for j := 1; j <= y; j++ {
				if j == 1 || j == y {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
		} else {
			for j := 1; j <= y; j++ {
				if j == 1 || j == y {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}

}
