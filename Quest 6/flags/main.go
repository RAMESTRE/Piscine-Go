package main

import (
	"fmt"
	"os"
)

func PrintMessage() {
	fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func Order(strings string) string {
	tab := []string{}
	var newString string
	for i := 0; i < len(strings); i++ {
		tab = append(tab, string(strings[i]))
	}
	for j := len(tab); j > 0; j-- {
		for k := 1; k < j; k++ {
			if tab[k] < tab[k-1] {
				tab[k], tab[k-1] = tab[k-1], tab[k]
			}
		}
	}
	for l := 0; l < len(tab); l++ {
		newString += tab[l]
	}
	return newString
}

func main() {
	var strings string
	var order bool

	if len(os.Args) < 2 {
		PrintMessage()
		return
	} else {
		for i := 1; i < len(os.Args); i++ {
			var verif string
			for j := 0; j < len(os.Args[i]); j++ {
				verif += string(os.Args[i][j])
				if verif == "--insert=" || verif == "-i=" {
					strings += os.Args[i][j+1:]
					break
				} else if verif == "--order" || verif == "-o" {
					order = true
				} else if verif == "--help" || verif == "-h" {
					PrintMessage()
					return
				} else if j == len(os.Args[i])-1 {
					strings = verif + strings
				}
			}
		}
		if order {
			strings = Order(strings)
		}
		fmt.Println(strings)
	}
}
