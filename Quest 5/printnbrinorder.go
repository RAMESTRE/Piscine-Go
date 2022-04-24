package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []int
	var save int
	if n == 0 {
		tab = append(tab, 0)
	}
	for n > 0 {
		save = n % 10
		tab = append(tab, save)
		n /= 10
	}
	for i := len(tab); i > 0; i-- {
		for j := 1; j < i; j++ {
			if tab[j] < tab[j-1] {
				tab[j], tab[j-1] = tab[j-1], tab[j]
			}
		}
	}
	for j := 0; j < len(tab); j++ {
		z01.PrintRune(rune(tab[j] + 48))
	}
}
