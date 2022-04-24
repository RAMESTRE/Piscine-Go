package tls_challenge_go_21_22

func TrimAtoi(s string) int {
	compteurNb := 0
	compteurMoins := 0
	var tab []int
	var nb int
	puissance := 1

	for i := 0; i < len(s); i++ {
		if compteurNb == 0 && compteurMoins == 0 && s[i] == 45 {
			compteurMoins += 1
		}
		if s[i] >= 48 && s[i] <= 57 {
			tab = append(tab, int(s[i]))
			compteurNb += 1
		}
	}
	for j := len(tab) - 1; j >= 0; j-- {
		nb = nb + (tab[j]-48)*puissance
		puissance *= 10
	}
	if compteurMoins == 1 && compteurNb > 0 {
		nb = -nb
	}
	if compteurNb == 0 {
		return 0
	}
	return nb
}
