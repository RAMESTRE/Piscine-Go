package tls_challenge_go_21_22

func IsPrintable(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] < 32 {
				return false
			}
		}
	}
	return true
}
