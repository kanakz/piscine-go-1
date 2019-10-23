package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ElementsCount(s []string) int {
	k := 0
	for index := range s {
		k = index
	}
	return k
}

func UpperCheck(s []string) bool {
	check := false
	if s[1] == "--upper" {
		check = true
	}
	return check
}

func main() {
	arguments := os.Args
	for i := 1; i <= ElementsCount(arguments); i++ {
		arg_str := arguments[i]
		arg_slice := []rune(arg_str)
		arg_asciinum := 0
		for _, letter := range arg_slice {
			if letter >= '0' && letter <= '9' {
				r := '0'
				if letter == r {
					arg_asciinum = arg_asciinum*10 + 0
				} else {
					k := 0
					for c := '0'; c < letter; c++ {
						k++
					}
					arg_asciinum = arg_asciinum*10 + k
				}
			} else {
				z01.PrintRune(' ')
				arg_asciinum = -1
				break
			}
		}
		if arg_asciinum == -1 {
			continue
		} else if arg_asciinum > 26 {
			z01.PrintRune(' ')
		} else {
			num := '@'
			for i := 1; i <= arg_asciinum; i++ {
				num++
			}
			if !UpperCheck(arguments) {
				num = num + 32
			}
			z01.PrintRune(num)
		}
	}
	z01.PrintRune('\n')
}

// package main

// import (
// 	"github.com/01-edu/z01"
// 	"os"
// )

// func main() {
// 	arguments := os.Args
// 	k := 0
// 	start := 1
// 	app := 'u'
// 	for index := range arguments {
// 		index = index
// 		k++
// 	}
// 	if k == 1 {
// 		k = k
// 	} else {
// 		if arguments[1] == "--upper" {
// 			start = 1
// 		} else {
// 			app = 'l'
// 		}
// 		for i := start; i <= k-1; i++ {
// 			str := arguments[i]
// 			slice := []rune(str)
// 			argInt := 0
// 			for _, letter := range slice {
// 				if letter >= '0' && letter <= '9' {
// 					a := '0'
// 					if letter == a {
// 						argInt = argInt*10 + 0
// 					} else {
// 						ck := 0
// 						for c := '0'; c < letter; c++ {
// 							ck++
// 						}
// 						argInt = argInt*10 + ck
// 					}
// 				} else {
// 					z01.PrintRune(' ')
// 					argInt = -1
// 					break
// 				}
// 			}
// 			if argInt == -1 {
// 				continue
// 			} else if argInt > 26 {
// 				z01.PrintRune(' ')
// 			} else {
// 				a := '@'
// 				for i := 1; i <= argInt; i++ {
// 					a++
// 				}
// 				if app == 'l' {
// 					a = a + 32
// 				}
// 				z01.PrintRune(a)
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func Atoi(s string) int {
// 	result := 0
// 	isNegative := 0
// 	for i, value := range s {

// 		if i == 0 && (value == '+' || value == '-') {
// 			if value == '-' {
// 				isNegative = 1
// 			}
// 			continue
// 		}
// 		count := 0
// 		if value < '0' || value > '9' {
// 			return 0
// 		}
// 		for i := '1'; i <= value; i++ {
// 			count++
// 		}
// 		result = result*10 + count
// 	}
// 	if isNegative == 1 {
// 		result = result * (-1)
// 	}
// 	return result
// }

// func ArrLen(arr []string) int {
// 	len := 0
// 	for i := range arr {
// 		len = i + 1
// 	}
// 	return len
// }

// func main() {
// 	args := os.Args
// 	if ArrLen(args) > 1 {
// 		count := 96
// 		if args[1] == "--upper" {
// 			count = 64
// 		}
// 		for i, str := range args {
// 			if i > 0 {
// 				nbr := Atoi(str) + count
// 				if nbr >= 97 && nbr <= 122 || nbr >= 65 && nbr <= 90 {
// 					z01.PrintRune(rune(nbr))
// 				} else {
// 					z01.PrintRune(' ')
// 				}

// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
