package tls_challenge_go_21_22

func IsSorted(f func(a, b int) int, a []int) bool {
	signe := 0
	for i := 1; i < len(a); i++ {
		if signe == 0 {
			if f(a[i-1], a[i]) < 0 {
				signe = -1
			} else if f(a[i-1], a[i]) > 0 {
				signe = 1
			}
		} else if signe == -1 {
			if f(a[i-1], a[i]) > 0 {
				return false
			}
		} else if signe == 1 {
			if f(a[i-1], a[i]) < 0 {
				return false
			}
		}
	}
	return true
}
