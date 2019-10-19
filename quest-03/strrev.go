package piscine

// func main() {
// 	input := "The quick brown 狐 jumped over the lazy 犬"
// 	// Get Unicode code points.
// 	n := 0
// 	rune := make([]rune, len(input))
// 	for _, r := range input {
// 			rune[n] = r
// 			n++
// 	}
// 	rune = rune[0:n]
// 	// Reverse
// 	for i := 0; i < n/2; i++ {
// 			rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
// 	}
// 	// Convert back to UTF-8.
// 	output := string(rune)
// 	fmt.Println(output)
// }

// func Reverse(s string) (result string) {
// 	for _,v := range s {
// 	  result = string(v) + result
// 	}
// 	return
// }

// func Reverse(s string) string {
//     runes := []rune(s)
//     for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }

func StrRev(s string) string {
	var ans string
	for _, c := range s {
		ans = string(c) + ans
	}
	return ans
}
