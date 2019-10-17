// package piscine

// import "github.com/01-edu/z01"

// func Nbr(n int) {
// 	a := '0'
// 	if n == 0 {
// 		z01.PrintRune(w)
// 		return
// 	}
// 	for k := 1; k <= n%10; k++ {
// 		a++
// 	}
// 	for k := -1; k >= n%10; k-- {
// 		a++
// 	}
// 	if n/10 != 0 {
// 		Nbr(n / 10)
// 	}
// 	z01.PrintRune(a)
// 	returnrintcombn
// }

// func check(p int) bool {
// 	for {
// 		if p == 0 {
// 			break
// 		}
// 		if p/10 != 0 && p%10 <= ((p/10)%10) {
// 			return false
// 		}
// 		p /= 10
// 	}
// 	return true
// }
// func ok(p int) bool {
// 	if p < 9 {
// 		return true
// 	}
// 	if p%10 == 9 {
// 		for {
// 			if p == 0 {
// 				break
// 			}
// 			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
// 				return true
// 			}
// 			p /= 10
// 		}

// 		return false
// 	} else {
// 		return true
// 	}
// }package printcombn

// import "github.com/01-edu/z01"

// func Nbr(n int) {
// 	a := '0'
// 	if n == 0 {
// 		z01.PrintRune(a)
// 		return
// 	}
// 	for k := 1; k <= n%10; k++ {
// 		a++
// 	}
// 	for k := -1; k >= n%10; k-- {
// 		a++
// 	}
// 	if n/10 != 0 {
// 		Nbr(n / 10)
// 	}
// 	z01.PrintRune(a)
// 	return
// }

// func check(p int) bool {
// 	for {
// 		if p == 0 {
// 			break
// 		}
// 		if p/10 != 0 && p%10 <= ((p/10)%10) {
// 			return false
// 		}
// 		p /= 10
// 	}
// 	return true
// }
// func ok(p int) bool {
// 	if p < 9 {
// 		return true
// 	}
// 	if p%10 == 9 {
// 		for {
// 			if p == 0 {
// 				break
// 			}
// 			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
// 				return true
// 			}
// 			p /= 10
// 		}

// 		return false
// 	} else {
// 		return true
// 	}
// }
// func PrintCombn(n int) {
// 	mx_ln := 1
// 	for i := 2; i <= n; i++ {
// 		mx_ln *= 10
// 	}
// 	for i := mx_ln / 10; i < mx_ln; i++ {
// 		if check(i) == true {
// 			if mx_ln >= 10 {
// 				z01.PrintRune('0')
// 			}
// 			Nbr(i)
// 			z01.PrintRune(',')
// 			z01.PrintRune(' ')
// 		}
// 	}
// 	for i := mx_ln; i <= mx_ln*9; i++ {
// 		if check(i) == true {
// 			Nbr(i)
// 			if ok(i) == true {
// 				z01.PrintRune(',')
// 				z01.PrintRune(' ')
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// func PrintCombn(n int) {
// 	mx_ln := 1
// 	for i := 2; i <= n; i++ {
// 		mx_ln *= 10
// 	}
// 	for i := mx_ln / 10; i < mx_ln; i++ {
// 		if check(i) == true {
// 			if mx_ln >= 10 {
// 				z01.PrintRune('0')
// 			}
// 			Nbr(i)
// 			z01.PrintRune(',')
// 			z01.PrintRune(' ')
// 		}
// 	}
// 	for i := mx_ln; i <= mx_ln*9; i++ {
// 		if check(i) == true {
// 			Nbr(i)
// 			if ok(i) == true {
// 				z01.PrintRune(',')
// 				z01.PrintRune(' ')
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// package piscine

// import (
// 	"fmt"
// )

// func show(n int, table [9]int, tmax [9]int) {
// 	i := 0
// 	for i < n {
// 		fmt.Print(table[i])
// 		i++
// 	}
// 	if table[0] != tmax[0] {
// 		fmt.Print(", ")
// 	}
// }

// func printComb1() {
// 	table := [9]int{0}
// 	tmax := [9]int{9}
// 	for table[0] <= tmax[0] {
// 		show(1, table, tmax)
// 		table[0]++
// 	}
// }

// func PrintCombN(n int) {
// 	table := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	tmax := [9]int{}

// 	if n == 1 {
// 		printComb1()
// 	} else {
// 		i := n - 1
// 		j := 9
// 		for i >= 0 {
// 			tmax[i] = j
// 			i--
// 			j--
// 		}
// 		i = n - 1
// 		for table[0] < tmax[0] {
// 			if table[i] != tmax[i] {
// 				show(n, table, tmax)
// 				table[i]++
// 			}
// 			if table[i] == tmax[i] {
// 				if table[i-1] != tmax[i-1] {
// 					show(n, table, tmax)
// 					table[i-1]++
// 					j = i
// 					for j < n {
// 						table[j] = table[j-1] + 1
// 						j++
// 					}
// 					i = n - 1
// 				}
// 			}
// 			for table[i] == tmax[i] && table[i-1] == tmax[i-1] && i > 1 {
// 				i--
// 			}
// 		}
// 		show(n, table, tmax)
// 	}
// 	fmt.Println()
// }
