package tls_challenge_go_21_22

func ToLower(s string) string {
	var newS string
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] >= 65 && s[i] <= 90 {
				newS += string(s[i] + 32)
			} else {
				newS += string(s[i])
			}
		}
	}
	return newS
}
