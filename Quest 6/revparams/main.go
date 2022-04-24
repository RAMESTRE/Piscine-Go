package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		argRune := []rune(os.Args[i])
		for j := 0; j < len(argRune); j++ {
			z01.PrintRune(argRune[j])
		}
		z01.PrintRune('\n')
	}
}
