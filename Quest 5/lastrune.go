package tls_challenge_go_21_22

func LastRune(s string) rune {
	sentence := []rune(s)
	lastChar := sentence[len(sentence)-1]
	return lastChar
}
