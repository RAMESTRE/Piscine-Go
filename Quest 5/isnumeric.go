package tls_challenge_go_21_22

func IsNumeric(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] < 48 || s[i] > 57 {
				return false
			}
		}
	}
	return true
}
