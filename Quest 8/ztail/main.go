package main

import (
	"fmt"
	"os"
)

func argOctetValable(arg string) int {
	var octet int
	runeArg := []rune(arg)
	puissance := 1
	for i := len(runeArg) - 1; i >= 0; i-- {
		if runeArg[i] < 48 || runeArg[i] > 57 {
			return -1
		}
		octet += (int(runeArg[i]) - 48) * puissance
		puissance *= 10
	}
	return octet - 1
}

func fileExist(arg string) []byte {
	readfile, err := os.ReadFile(arg)
	if err != nil {
		fmt.Printf("open %s: no such file or directory\n", arg)
		return readfile
	}
	return readfile
}

func printOctetDemandes(tab []byte, octet int) {
	var texte string
	for i := len(tab) - octet; i < len(tab); i++ {
		texte += string(tab[i])
	}
	fmt.Printf(texte + "\n")
}

func deleteNewLine(readfile []byte) []byte {
	var tab []byte
	for len(readfile) > 0 {
		if readfile[0] != 10 {
			tab = append(tab, readfile[0])
		}
		readfile = readfile[1:]
	}
	return tab
}

func main() {
	os.Args = os.Args[1:]
	var count int
	var countSuccessfullFile int
	if len(os.Args) >= 2 {
		octet := argOctetValable(os.Args[1])
		if octet >= 0 && os.Args[0] == "-c" {
			os.Args = os.Args[2:]
			for i := 0; i < len(os.Args); i++ {
				tab := fileExist(os.Args[i])
				tabOctet := deleteNewLine(tab)
				if tabOctet != nil {
					if count > 0 {
						fmt.Printf("\n")
					}
					if len(os.Args) > 1 {
						fmt.Printf("==> %s <==\n", os.Args[i])
					}
					if octet > len(tabOctet) {
						octet = len(tabOctet)
						printOctetDemandes(tabOctet, octet)
						os.Exit(1)
					}
					printOctetDemandes(tabOctet, octet)
					countSuccessfullFile += 1
				}
				count += 1
			}
		}
		if countSuccessfullFile == len(os.Args) {
			return
		}
	}
	os.Exit(1)
}
