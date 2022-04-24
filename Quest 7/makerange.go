package tls_challenge_go_21_22

func MakeRange(min, max int) []int {
	lenght := max - min
	tab := []int{}

	if min >= max {
		tab = nil
	} else {
		tab = make([]int, lenght)
		for i := 0; i < lenght; i++ {
			tab[i] = min + i
		}
	}
	return tab
}
