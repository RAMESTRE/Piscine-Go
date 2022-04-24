package tls_challenge_go_21_22

func SplitWhiteSpaces(s string) []string {
	letterWord := 0
	var tab []string
	var word string

	for i := 0; i < len(s); i++ {
		if s[i] == 32 || s[i] == 9 || s[i] == 10 {
			if letterWord > 0 {
				tab = append(tab, word)
			}
			letterWord = 0
			word = ""
		} else {
			letterWord += 1
			word += string(s[i])
		}
		if i == len(s)-1 {
			tab = append(tab, word)
		}
	}
	return tab
}
