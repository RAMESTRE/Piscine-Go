package tls_challenge_go_21_22

func CountIf(f func(string) bool, a []string) int {
	count := 0
	for _, element := range a {
		resultat := f(element)
		if resultat == true {
			count += 1
		}
	}
	return count
}
