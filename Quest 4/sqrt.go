package tls_challenge_go_21_22

func Sqrt(nb int) int {
	for i := nb; i > 0; i-- {
		var root int
		root = i * i
		if root == nb {
			return i
		}
	}
	return 0
}
