package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Printparams() {
	for k := 1; k < len(os.Args); k++ {
		argRune := []rune(os.Args[k])
		for l := 0; l < len(argRune); l++ {
			z01.PrintRune(argRune[l])
		}
		z01.PrintRune('\n')
	}
}

func main() {
	for i := len(os.Args); i > 0; i-- {
		for j := 2; j < i; j++ {
			if os.Args[j] < os.Args[j-1] {
				os.Args[j], os.Args[j-1] = os.Args[j-1], os.Args[j]
			}
		}
	}
	Printparams()
}
