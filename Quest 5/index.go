package tls_challenge_go_21_22

func Index(s string, toFind string) int {
	comp := 0
	if len(s) > 0 && len(toFind) > 0 {
		for i := 0; i < len(s); i++ {
			if toFind[0] == s[i] {
				comp = i
				break
			}
		}
		for j := 0; j < len(toFind); j++ {
			if toFind[j] != s[comp+j] {
				return -1
			}
		}
	} else if len(toFind) == 0 {
		return 0
	}
	return comp
}
