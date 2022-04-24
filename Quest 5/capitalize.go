package tls_challenge_go_21_22

func Capitalize(s string) string {
	lettre := 0
	var newS string
	if len(s) == 0 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if lettre == 0 {
			if s[i] >= 97 && s[i] <= 122 {
				newS += string((s[i]) - 32)
			} else {
				newS += string(s[i])
			}
		} else if lettre > 0 {
			if s[i] >= 65 && s[i] <= 90 {
				newS += string((s[i]) + 32)
			} else {
				newS += string(s[i])
			}
		}
		lettre += 1
		if s[i] < 48 || s[i] > 57 && s[i] < 65 || s[i] > 90 && s[i] < 97 || s[i] > 122 {
			lettre = 0
		}
	}
	return newS
}
