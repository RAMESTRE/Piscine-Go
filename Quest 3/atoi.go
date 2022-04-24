package tls_challenge_go_21_22

func Atoi(s string) int {
	nb := 0
	puissance := 1
	negatif := false
	if s == "" {
		return 0
	}
	if s[0] == 45 {
		negatif = true
		s = s[1:]
	} else if s[0] == 43 {
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < 48 || s[i] > 57 {
			return 0
		} else {
			nb = nb + (int(s[i])-48)*puissance
			puissance = puissance * 10
		}
	}
	if negatif == true {
		nb = -nb
	}
	return nb
}
