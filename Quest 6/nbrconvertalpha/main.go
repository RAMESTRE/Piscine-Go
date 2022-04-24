package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var letter int

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if len(os.Args[i]) == 1 {
				letter = int(os.Args[i][0]) - 48 + 96
			} else if len(os.Args[i]) == 2 {
				letter = (int(os.Args[i][0])-48)*10 + int(os.Args[i][1]) - 48 + 96
			}
			if os.Args[1] == "--upper" {
				letter = letter - 32
			}
			if letter < 65 || letter > 90 && letter < 97 || letter > 122 || len(os.Args[i]) > 2 {
				if os.Args[i] != "--upper" {
					z01.PrintRune(' ')
				}
			} else {
				z01.PrintRune(rune(letter))
			}
		}
		z01.PrintRune('\n')
	}
}
