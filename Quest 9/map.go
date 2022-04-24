package tls_challenge_go_21_22

func Map(f func(int) bool, a []int) []bool {
	var tabBool []bool
	for _, element := range a {
		tabBool = append(tabBool, f(element))
	}
	return tabBool
}
