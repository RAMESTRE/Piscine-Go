package tls_challenge_go_21_22

func StrRev(s string) string {
	var reversedString string
	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}
	return reversedString
}
