package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintStr(s string) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
}
