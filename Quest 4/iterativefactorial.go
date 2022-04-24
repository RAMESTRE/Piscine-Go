package tls_challenge_go_21_22

func IterativeFactorial(nb int) int {
	compteur := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		compteur *= i
	}
	return compteur
}
