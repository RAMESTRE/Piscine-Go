package tls_challenge_go_21_22

func ValidBase(base string) bool {
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

func NombreInBase(nb string, base string) bool {
	for i := 0; i < len(base); i++ {
		if string(base[i]) == nb {
			return true
		}
	}
	return false
}

func IndexInBase(nb string, base string) int {
	for i := 0; i < len(base); i++ {
		if string(base[i]) == nb {
			return i
		}
	}
	return 0
}

func ConvertToBaseTen(nbr, baseFrom string) int {
	base := len(baseFrom)
	puissance := 1
	var nbBaseTen int
	for i := len(nbr) - 1; i >= 0; i-- {
		if NombreInBase(string(nbr[i]), baseFrom) {
			nbBaseTen += IndexInBase(string(nbr[i]), baseFrom) * puissance
			puissance *= base
		}
	}
	return nbBaseTen
}

func ConvertToBaseN(nbrBaseTen int, baseTo string) string {
	baseN := len(baseTo)
	var reste int
	var nbrBaseN string
	var tab []int
	if nbrBaseTen == 0 {
		return string(baseTo[0])
	}
	for nbrBaseTen > 0 {
		reste = nbrBaseTen % baseN
		tab = append(tab, reste)
		nbrBaseTen /= baseN
	}
	for i := len(tab) - 1; i >= 0; i-- {
		nbrBaseN += string(baseTo[tab[i]])
	}
	return nbrBaseN
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var nbrBaseTen int
	if ValidBase(baseFrom) {
		for i := 0; i < len(nbr); i++ {
			bool := NombreInBase(string(nbr[i]), baseFrom)
			if bool == false {
				return "0"
			}
		}
		nbrBaseTen = ConvertToBaseTen(nbr, baseFrom)
	} else {
		return "0"
	}
	return ConvertToBaseN(nbrBaseTen, baseTo)
}
