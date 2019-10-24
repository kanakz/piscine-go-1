package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	k := 0
	toFind1 := "--insert="
	toFind2 := "-i="
	isOrder := 0
	isOrderPos := 1
	isInsert := 0
	str := ""
	q := 0
	strInsert := ""
	var slice []rune
	for index := range arguments {
		if arguments[index] == "--order" || arguments[index] == "-o" {
			isOrder = 1
			isOrderPos = index
		}
		k++
	}
	if k == 1 {
		help()
		return
	} else if arguments[1] == "--help" || arguments[1] == "-h" {
		help()
		return
	} else if k == 2 {
		fmt.Println(arguments[1])
	} else {
		//opredelyaem est li insert i formireum slice iz pervogo elementa

		if Index(arguments[1], toFind1) == 0 {
			isInsert = 1
			slice1 := []rune(arguments[1])
			for index, letter := range slice1 {
				if index >= 9 {
					slice = append(slice, letter)
				}
			}
		}
		if Index(arguments[1], toFind2) == 0 {
			isInsert = 1
			slice1 := []rune(arguments[1])
			for index, letter := range slice1 {
				if index >= 3 {
					slice = append(slice, letter)
				}
			}
		}

		strInsert = string(slice)
		//snachala formiruem samu stroku, esli ona razbita na chasti
		if isOrder == 1 {
			q = isOrderPos + 1
		} else if isInsert == 1 {
			q = 2
		} else {
			q = 1
		}
		str = arguments[q]
		for i := q + 1; i < k; i++ {
			str = str + " " + arguments[i]
		}
		//teper ishem stroku inserta, esli ona razbita na chasti
		if isInsert == 1 {
			if isOrder == 1 && isOrderPos > 2 {
				for i := 2; i < isOrderPos; i++ {
					strInsert = strInsert + " " + arguments[i]
				}
			}
			str = str + strInsert
		}
		if isOrder == 1 {
			slice = []rune(str)
			for index1 := range slice {
				if index1 > 0 {
					for index, letter := range slice {
						if index > 0 {
							if letter < slice[index-1] {
								slice[index] = slice[index-1]
								slice[index-1] = letter
							}
						}
					}
				}
			}
		} else {
			slice = []rune(str)
		}
		for _, letter := range slice {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func Index(s, toFind string) int {
	sliceS := []rune(s)
	sliceF := []rune(toFind)
	k := 0
	for index := range sliceF {
		index = index
		k++
	}
	q := 0
	for index := range sliceS {
		index = index
		q++
	}
	if k > q {
		return -1
	}
	for index, letter := range sliceS {
		if letter == sliceF[0] && q >= k+index-1 {
			m := 1
			for i := 1; i < k; i++ {
				if sliceF[i] == sliceS[index+i] {
					m++
				}
			}
			if m == k {
				return index
			}
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func Show() {
// 	fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
// }
// func Sort(s string) {
// 	var a [5000]int
// 	for _, c := range s {
// 		a[int(c)]++
// 	}
// 	for i, c := range a {
// 		for c > 0 {
// 			z01.PrintRune(rune(i))
// 			c--
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// func main() {
// 	arg := os.Args[1:]
// 	ans := ""
// 	ok := false
// 	SortIt := false
// 	for _, c := range arg {
// 		ok = true
// 		if c == "-h" || c == "--help" {
// 			Show()
// 			break

// 		}
// 		ln := 0
// 		for j := range c {
// 			ln = j + 1
// 		}
// 		if ln > 0 {
// 			if c[0] == '-' {
// 				if ln > 2 && c[2] == 'i' {
// 					if ln > 8 {
// 						ans += c[9:]
// 					}
// 				}
// 				if c[1] == 'i' {
// 					if ln > 3 {
// 						ans += c[3:]
// 					}
// 				}
// 			} else {
// 				ans = c + ans
// 			}
// 		}
// 		if c == "-o" || c == "--order" {
// 			SortIt = true
// 		}
// 	}
// 	if !ok {
// 		Show()
// 	}
// 	if SortIt {
// 		Sort(ans)
// 	} else {
// 		fmt.Println(ans)
// 	}
// }
