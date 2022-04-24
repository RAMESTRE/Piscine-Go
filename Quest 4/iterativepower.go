package tls_challenge_go_21_22

func IterativePower(nb int, power int) int {
	mult := nb
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 1; i < power; i++ {
		nb *= mult
	}
	return nb
}
