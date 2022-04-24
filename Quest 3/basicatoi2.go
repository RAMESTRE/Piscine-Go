package tls_challenge_go_21_22

func BasicAtoi2(s string) int {
	nb := 0
	puissance := 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < 48 || s[i] > 57 {
			return 0
		} else {
			nb = nb + (int(s[i])-48)*puissance
			puissance = puissance * 10
		}
	}
	return nb
}
