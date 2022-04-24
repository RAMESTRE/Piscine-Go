package tls_challenge_go_21_22

func RecursiveFactorial(arg int) int {
	if arg < 0 || arg > 20 {
		return 0
	} else if arg == 0 {
		return 1
	}
	return arg * RecursiveFactorial(arg-1)
}
