package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptrx, ptry *int) {
	*ptrx = 42
	*ptry = 21
}

func printNb(nb int) {
	var reste int
	var tab []int

	for nb > 0 {
		reste = nb % 10
		tab = append(tab, reste)
		nb /= 10
	}
	for i := len(tab) - 1; i >= 0; i-- {
		z01.PrintRune(rune(int(tab[i]) + 48))
	}
}

func main() {
	var x int
	var y int

	setPoint(&x, &y)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNb(x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printNb(y)
	z01.PrintRune('\n')
}
