package tls_challenge_go_21_22

func AppendRange(min, max int) []int {
	tab := []int{}
	if min >= max {
		tab = nil
	} else {
		for ; min < max; min++ {
			tab = append(tab, min)
		}
	}
	return tab
}
