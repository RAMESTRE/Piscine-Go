package tls_challenge_go_21_22

func IsPrime(nb int) bool {
	PrimeNb := nb
	if nb < 2 {
		return false
	}
	for i := 2; i < nb/2; i++ {
		PrimeNb = nb % i
		if PrimeNb == 0 {
			return false
		}
	}
	return true
}
