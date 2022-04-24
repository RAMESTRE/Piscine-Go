package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var save int
	var tableau []int
	var conversion int

	if n < 0 {
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		save = n % 10
		if save < 0 {
			save = -save
		}
		tableau = append(tableau, save)
		n /= 10
	}
	for i := len(tableau) - 1; i >= 0; i-- {
		conversion = tableau[i] + 48
		z01.PrintRune(rune(conversion))
	}
}
