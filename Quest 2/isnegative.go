package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func IsNegative(n int) {
	if n >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')
}
