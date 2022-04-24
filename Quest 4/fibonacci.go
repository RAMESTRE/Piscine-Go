package tls_challenge_go_21_22

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else if index <= 0 {
		return -1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
