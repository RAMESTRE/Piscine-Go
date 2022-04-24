package tls_challenge_go_21_22

func SortWordArr(a []string) []string {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	return a
}
