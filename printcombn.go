package piscine

import "github.com/01-edu/z01"

func Nbr(n int) {
	a := '0'
	if n == 0 {
		z01.PrintRune(w)
		return
	}
	for k := 1; k <= n%10; k++ {
		a++
	}
	for k := -1; k >= n%10; k-- {
		a++
	}
	if n/10 != 0 {
		Nbr(n / 10)
	}
	z01.PrintRune(a)
	returnrintcombn
}

func check(p int) bool {
	for {
		if p == 0 {
			break
		}
		if p/10 != 0 && p%10 <= ((p/10)%10) {
			return false
		}
		p /= 10
	}
	return true
}
func ok(p int) bool {
	if p < 9 {
		return true
	}
	if p%10 == 9 {
		for {
			if p == 0 {
				break
			}
			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
				return true
			}
			p /= 10
		}

		return false
	} else {
		return true
	}
}package printcombn

import "github.com/01-edu/z01"

func Nbr(n int) {
	a := '0'
	if n == 0 {
		z01.PrintRune(a)
		return
	}
	for k := 1; k <= n%10; k++ {
		a++
	}
	for k := -1; k >= n%10; k-- {
		a++
	}
	if n/10 != 0 {
		Nbr(n / 10)
	}
	z01.PrintRune(a)
	return
}

func check(p int) bool {
	for {
		if p == 0 {
			break
		}
		if p/10 != 0 && p%10 <= ((p/10)%10) {
			return false
		}
		p /= 10
	}
	return true
}
func ok(p int) bool {
	if p < 9 {
		return true
	}
	if p%10 == 9 {
		for {
			if p == 0 {
				break
			}
			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
				return true
			}
			p /= 10
		}

		return false
	} else {
		return true
	}
}
func PrintCombn(n int) {
	mx_ln := 1
	for i := 2; i <= n; i++ {
		mx_ln *= 10
	}
	for i := mx_ln / 10; i < mx_ln; i++ {
		if check(i) == true {
			if mx_ln >= 10 {
				z01.PrintRune('0')
			}
			Nbr(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for i := mx_ln; i <= mx_ln*9; i++ {
		if check(i) == true {
			Nbr(i)
			if ok(i) == true {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
func PrintCombn(n int) {
	mx_ln := 1
	for i := 2; i <= n; i++ {
		mx_ln *= 10
	}
	for i := mx_ln / 10; i < mx_ln; i++ {
		if check(i) == true {
			if mx_ln >= 10 {
				z01.PrintRune('0')
			}
			Nbr(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for i := mx_ln; i <= mx_ln*9; i++ {
		if check(i) == true {
			Nbr(i)
			if ok(i) == true {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}