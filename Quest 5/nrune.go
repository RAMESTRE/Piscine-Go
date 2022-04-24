package tls_challenge_go_21_22

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		sentence := []rune(s)
		nChar := sentence[n-1]
		return nChar
	}
	return 0
}
