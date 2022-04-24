package main

import (
	"os"

	"github.com/01-edu/z01"
)

func RotateVowels(newString string, tabIndex []int) {
	tab := []string{}
	var stringFinal string

	for i := 0; i < len(newString); i++ {
		tab = append(tab, string(newString[i]))
	}
	for j := 0; j < len(tabIndex)/2; j++ {
		tab[tabIndex[j]], tab[tabIndex[len(tabIndex)-j-1]] = tab[tabIndex[len(tabIndex)-j-1]], tab[tabIndex[j]]
	}
	for k := 0; k < len(tab); k++ {
		stringFinal += tab[k]
	}
	runeFinal := []rune(stringFinal)

	for l := 0; l < len(runeFinal); l++ {
		z01.PrintRune(runeFinal[l])
	}
}

func main() {
	lenghtArg := len(os.Args)
	var newString string
	var tabIndex []int

	if lenghtArg < 2 {
		z01.PrintRune('\n')
		return
	}
	for i := 1; i < lenghtArg; i++ {
		newString += os.Args[i]
		newString += " "
	}
	for j := 0; j < len(newString); j++ {
		if newString[j] == 97 || newString[j] == 101 || newString[j] == 105 || newString[j] == 111 || newString[j] == 117 || newString[j] == 65 || newString[j] == 69 || newString[j] == 73 || newString[j] == 79 || newString[j] == 85 {
			tabIndex = append(tabIndex, j)
		}
	}
	RotateVowels(newString, tabIndex)
	z01.PrintRune('\n')
}
