package tls_challenge_go_21_22

import "github.com/01-edu/z01"

func PrintWord(word string) {
	runeWord := []rune(word)
	for i := 0; i < len(runeWord); i++ {
		z01.PrintRune(runeWord[i])
	}
	z01.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	if len(a) > 0 {
		for i := 0; i < len(a); i++ {
			PrintWord(a[i])
		}
	}
}
