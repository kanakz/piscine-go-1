package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			if j == y || j == 1 {
				for i := 1; i <= x; i++ {
					if i == 1 || i == x {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				}
			} else {
				for i := 1; i <= x; i++ {
					if i == 1 || i == x {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}
