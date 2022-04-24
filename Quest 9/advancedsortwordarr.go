package tls_challenge_go_21_22

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			result := f(a[i], a[j])
			if result == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
