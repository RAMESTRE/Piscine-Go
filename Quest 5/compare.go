package tls_challenge_go_21_22

func Compare(a, b string) int {
	tailleA := len(a)
	tailleB := len(b)
	var tour int
	if tailleA > tailleB {
		tour = tailleB
	} else {
		tour = tailleA
	}
	for i := 0; i < tour; i++ {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}
