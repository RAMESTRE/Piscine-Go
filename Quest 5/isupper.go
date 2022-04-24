package tls_challenge_go_21_22

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 65 || s[i] > 90 {
			return false
		}
	}
	return true
}
