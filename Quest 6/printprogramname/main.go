package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argRune := []rune(os.Args[0])

	for i := 0; i < len(argRune); i++ {
		if argRune[i] != '/' && argRune[i] != '.' {
			z01.PrintRune(argRune[i])
		}
	}
	z01.PrintRune('\n')
}
