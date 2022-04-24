package tls_challenge_go_21_22

func IsAlpha(s string) bool {
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] < 48 || s[i] > 57 && s[i] < 65 || s[i] > 90 && s[i] < 97 || s[i] > 122 {
				return false
			}
		}
	}
	return true
}
