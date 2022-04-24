package tls_challenge_go_21_22

import (
	"github.com/01-edu/z01"
)

func verifCase(tab []int, i int) bool {
	for j := 0; j < i; j++ {
		ligneDiag := i - j
		colonneDiag := tab[i] - tab[j]
		if ligneDiag < 0 {
			ligneDiag = -ligneDiag
		} else if colonneDiag < 0 {
			colonneDiag = -colonneDiag
		}
		if ligneDiag-colonneDiag == 0 {
			return false
		} else if tab[i] == tab[j] {
			return false
		}
	}
	return true
}

func printz01(tab []int) {
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(rune(tab[i] + 48))
	}
	z01.PrintRune('\n')
}

func resolution(tab []int, n int, compteur int) {
	if compteur < n {
		for i := 1; i <= n; i++ {
			tab[compteur] = i
			if verifCase(tab, compteur) == true {
				if compteur == n-1 {
					printz01(tab)
				}
				resolution(tab, n, compteur+1)
			}
		}
	}
}

func EightQueens() {
	n := 8
	compteur := 0
	tab := []int{0, 0, 0, 0, 0, 0, 0, 0}
	resolution(tab, n, compteur)
}
