package tls_challenge_go_21_22

func Join(strs []string, sep string) string {
	var sentence string
	for i := 0; i < len(strs); i++ {
		sentence += strs[i]
		if i < len(strs)-1 {
			sentence += sep
		}
	}
	return sentence
}
