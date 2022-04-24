package tls_challenge_go_21_22

func SortIntegerTable(table []int) {
	for i := len(table); i > 0; i-- {
		for j := 1; j < i; j++ {
			if table[j] < table[j-1] {
				table[j], table[j-1] = table[j-1], table[j]
			}
		}
	}
}
