package piscine

func MinusPos(s string) bool {
	minus := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			break
		}
		if c == '-' {
			minus = true
		}
	}
	return minus
}

func StrCheck(s string) bool {
	check := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			check = true
			break
		}
	}
	return check
}

func TrimAtoi(s string) int {
	x := 0
	k := 0
	if StrCheck(s) {
		for _, c := range s {
			if c >= '0' && c <= '9' {
				for i := '1'; i <= c; i++ {
					k++
				}
				x = x*10 + k
			}
		}
	}
	if MinusPos(s) {
		x *= -1
	}
	return x
}
