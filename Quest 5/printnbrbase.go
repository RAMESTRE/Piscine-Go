package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	baseN := len(base)
	var tab []int
	var reste int
	valide := true
	baseRune := []rune(base)

	for i := 0; i < len(base); i++ {
		if string(base[i]) == "-" || string(base[i]) == "+" {
			valide = false
			break
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				valide = false
				break
			}
		}
	}
	if len(base) > 1 && valide != false {
		if nbr == 0 {
			z01.PrintRune(baseRune[0])
		}
		if nbr < 0 {
			z01.PrintRune('-')
		}
		for nbr < 0 {
			reste = nbr % baseN
			tab = append(tab, reste)
			nbr /= baseN
		}
		for nbr > 0 {
			reste = nbr % baseN
			tab = append(tab, reste)
			nbr /= baseN
		}
		for i := len(tab) - 1; i >= 0; i-- {
			if tab[i] < 0 {
				tab[i] = -tab[i]
			}
			z01.PrintRune(baseRune[tab[i]])
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')
	}
}
