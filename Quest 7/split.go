package tls_challenge_go_21_22

func Split(s, sep string) []string {
	var word string
	countLetter := 0
	tab := make([]string, 0)

	for i := 0; i < len(s); i++ {
		count := 0
		for j := 0; j < len(sep); j++ {
			if i+j < len(s) && sep[j] == s[i+j] {
				count += 1
			}
		}
		if count != len(sep) {
			word += string(s[i])
			countLetter += 1
		} else {
			tab = append(tab, word)
			word = ""
			i = i + len(sep) - 1
			countLetter = 0
		}
		if i == len(s)-1 && countLetter > 0 {
			tab = append(tab, word)
		}
	}
	return tab
}
