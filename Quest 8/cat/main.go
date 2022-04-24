package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printTextFile(readfile string) {
	runeReadfile := []rune(readfile)
	for i := 0; i < len(runeReadfile); i++ {
		z01.PrintRune(runeReadfile[i])
	}
}

func printError(args string) {
	err := "ERROR: open "

	for i := 0; i < len(args); i++ {
		err += string(args[i])
	}

	err += ": no such file or directory"

	runeError := []rune(err)

	for j := 0; j < len(runeError); j++ {
		z01.PrintRune(runeError[j])
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			readFile, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				printError(os.Args[i])
				os.Exit(1)
			}
			printTextFile(string(readFile))
		}
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return
		}
		printTextFile(string(bytes))
	}
}
