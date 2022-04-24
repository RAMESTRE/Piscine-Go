package tls_challenge_go_21_22

func Any(f func(string) bool, a []string) bool {
	for _, element := range a {
		resultat := f(element)
		if resultat == true {
			return true
		}
	}
	return false
}
