package piscine

import "fmt"

const n = 8

var cells = [n]int{}

func PositionCheck(queen, pos int) bool {
	for i := 0; i < queen; i++ {
		_pos := cells[i]
		if _pos == pos || _pos == pos-(queen-i) || _pos == pos+(queen-i) {
			return false
		}
	}
	return true
}

func SetQueens(k int) {
	if k == n {
		for i := 0; i < n; i++ {
			fmt.Print(cells[i] + 1)
		}
		fmt.Println('\n')
	} else {
		for i := 0; i < n; i++ {
			if PositionCheck(k, i) {
				cells[k] = i
				SetQueens(k + 1)
			}
		}
	}
}

func EightQueens() {
	SetQueens(0)
}
