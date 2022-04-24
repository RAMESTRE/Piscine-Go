package tls_challenge_go_21_22

func BasicJoin(elems []string) string {
	var sentence string
	for i := 0; i < len(elems); i++ {
		sentence += elems[i]
	}
	return sentence
}
