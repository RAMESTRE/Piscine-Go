package tls_challenge_go_21_22

func AlphaCount(s string) int {
	compteur := 0
	sentence := []rune(s)
	for i := 0; i < len(s); i++ {
		if sentence[i] > 64 && sentence[i] < 91 || sentence[i] > 96 && sentence[i] < 123 {
			compteur += 1
		}
	}
	return compteur
}
