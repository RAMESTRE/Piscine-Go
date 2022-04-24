package tls_challenge_go_21_22

func ToUpper(s string) string {
	var newS string
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] >= 97 && s[i] <= 122 {
				newS += string(s[i] - 32)
			} else {
				newS += string(s[i])
			}
		}
	}
	return newS
}
