package tls_challenge_go_21_22

func FindNextPrime(nb int) int {
	PrimeNb := nb
	if nb < 2 {
		return 2
	}
	for j := 2; j <= nb/2; j++ {
		PrimeNb = nb % j
		if PrimeNb == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
