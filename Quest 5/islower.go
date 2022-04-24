package tls_challenge_go_21_22

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 97 || s[i] > 122 {
			return false
		}
	}
	return true
}
