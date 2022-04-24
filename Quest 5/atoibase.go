package tls_challenge_go_21_22

func BaseCorrecte(base string) bool {
	for i := 0; i < len(base); i++ {
		if string(base[i]) == "-" || string(base[i]) == "+" {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func NombreDsBase(nb string, base string) bool {
	for i := 0; i < len(base); i++ {
		if string(base[i]) == nb {
			return true
		}
	}
	return false
}

func IndexDsBase(nb string, base string) int {
	for i := 0; i < len(base); i++ {
		if string(base[i]) == nb {
			return i
		}
	}
	return 0
}

func AtoiBase(s string, base string) int {
	puissance := 1
	var tab []int
	baseN := len(base)
	var nombreBaseDix int
	verif := BaseCorrecte(base)

	if verif == false {
		return 0
	}
	for i := len(s) - 1; i >= 0; i-- {
		if NombreDsBase(string(s[i]), base) == true {
			tab = append(tab, IndexDsBase(string(s[i]), base))
		}
	}
	for i := 0; i < len(tab); i++ {
		nombreBaseDix += tab[i] * puissance
		puissance *= baseN
	}
	return nombreBaseDix
}
