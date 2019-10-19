// package main

// import (
// 	"fmt"

// 	piscine ".."
// )

// func main() {
// 	s := []int{5, 4, 3, 2, 1, 0, 100, 10}
// 	piscine.SortIntegerTable(s)
// 	fmt.Println(s)
	// s := "12345"
	// s2 := "0000000012345"
	// s3 := "012 345"
	// s4 := "Hello World!"
	// s5 := "+1234"
	// s6 := "-1234"
	// s7 := "++1234"
	// s8 := "--1234"

	// n := piscine.Atoi(s)
	// n2 := piscine.Atoi(s2)
	// n3 := piscine.Atoi(s3)
	// n4 := piscine.Atoi(s4)
	// n5 := piscine.Atoi(s5)
	// n6 := piscine.Atoi(s6)
	// n7 := piscine.Atoi(s7)
	// n8 := piscine.Atoi(s8)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)
	// fmt.Println(n5)
	// fmt.Println(n6)
	// fmt.Println(n7)
	// fmt.Println(n8)
	// s := "12345"
	// s2 := "0000000012345"
	// s3 := "012 345"
	// s4 := "Hello World!"

	// n := piscine.BasicAtoi2(s)
	// n2 := piscine.BasicAtoi2(s2)
	// n3 := piscine.BasicAtoi2(s3)
	// n4 := piscine.BasicAtoi2(s4)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)
	// s := "12345"
	// s2 := "0000000012345"
	// s3 := "000000"

	// n := piscine.BasicAtoi(s)
	// n2 := piscine.BasicAtoi(s2)
	// n3 := piscine.BasicAtoi(s3)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// s := "Hello World!"
	// s = piscine.StrRev(s)
	// fmt.Println(s)
	// a := 0
	// b := 1
	// piscine.Swap(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// str := "Hello World!"
	// nb := piscine.StrLen(str)
	// fmt.Println(nb)
	// str := "Hello World!"
	// piscine.PrintStr(str)
	// a := 13
	// b := 2
	// piscine.UltimateDivMod(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// a := 13
	// b := 2
	// var div int
	// var mod int
	// piscine.DivMod(a, b, &div, &mod)
	// fmt.Println(div)
	// fmt.Println(mod)
	// a := 0
	// b := &a
	// n := &b
	// piscine.UltimatePointOne(&n)
	// fmt.Println(a)
	// n := 0
	// piscine.PointOne(&n)
	// fmt.Println(n)
}

// package main

// func addTen(b *int) {
// 	*b = *b + 10
// }

// func main() {
// 	a := 10
// 	addTen(&a)
// 	println(a)
// }

// func addTen(b int) int {
// 	b = b + 10
// 	return b
// }

// func main() {
// 	a := 10
// 	a = addTen(a)

// 	println(a)

// a := 10

// fmt.Println(a)
// fmt.Printf("This is the adress of a : %v", &a)
// fmt.Println()

// var pointer *int = &a

// fmt.Println(pointer)

// slice := []string{
// 	"Word1",
// 	"Word2",
// 	"Word3",
// 	"Word4",
// 	"Word5",
// 	"Word6",
// }

// for index, word := range slice {

// 	fmt.Printf("The index is: %v the element is: %v\n", index, word)
// }

// for _, word := range slice {
// 	fmt.Printf("the element is: %v\n", word)
// }

// aString := "Hello"

// aStringChangeable := []byte(aString)

// aStringChangeable[0] = 'A'

// aStringFinalized := string(aStringChangeable)

// fmt.Println(aString) Hello

// fmt.Println(aStringChangeable) [65 101 108 108 111]

// fmt.Println(aStringFinalized) Aello
//}
