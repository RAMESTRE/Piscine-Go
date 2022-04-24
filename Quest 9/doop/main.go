package main

import (
	"os"
)

func resultValid(arg string, a, b int) bool {
	var result uint
	negatif := false
	if a < 0 {
		a = -a
	} else if b < 0 {
		b = -b
	}
	if arg == "+" {
		result = uint(a) + uint(b)
		if a < 0 && b < 0 {
			negatif = true
		}
	} else if arg == "-" {
		result = uint(a) + uint(b)
		if a < 0 && b < 0 {
			negatif = true
		}
	} else if arg == "*" {
		var max int
		var min int
		if a > b {
			max = a
			min = b
		} else if a < b {
			max = b
			min = a
		}
		for i := 0; i < max; i++ {
			result += uint(a) * uint(min)
			if result > 9223372036854775807 || result > 9223372036854775808 && negatif == true {
				return false
			}
		}
		if a < 0 || b < 0 {
			negatif = true
		}
	}
	if result > 9223372036854775807 || result > 9223372036854775808 && negatif == true {
		return false
	}
	return true
}

func operation(arg string, a, b int) string {
	var result int
	var newString string
	var save int
	var signeString string

	if arg == "+" {
		if resultValid(arg, a, b) == false {
			return newString
		}
		result = a + b
	} else if arg == "-" {
		if resultValid(arg, a, b) == false {
			return newString
		}
		result = a - b
	} else if arg == "*" {
		if resultValid(arg, a, b) == false {
			return newString
		}
		result = a * b
	} else if arg == "/" {
		if b == 0 {
			os.Stderr.WriteString("No division by 0\n")
			return newString
		} else {
			result = a / b
		}
	} else if arg == "%" {
		if b == 0 {
			os.Stderr.WriteString("No modulo by 0\n")
			return newString
		} else {
			result = a % b
		}
	}

	if result == 0 {
		return "0"
	}
	if result < 0 {
		signeString = "-"
		result = -result
	}
	for result > 0 {
		save = result % 10
		newString = string(save+48) + newString
		result /= 10
	}
	if signeString != "" {
		newString = signeString + newString
	}
	return newString
}

func convertStringNumber(arg string) int {
	puissance := 1
	var nb int
	negatif := false

	for i := len(arg) - 1; i >= 0; i-- {
		if i == 0 && arg[i] == 45 {
			negatif = true
		} else {
			nb += int(arg[i]-48) * puissance
			puissance *= 10
		}
	}
	if negatif == true {
		return -nb
	}
	return nb
}

func validValue(arg string) bool {
	puissance := 1
	negatif := false
	argSave := arg

	var nb uint
	for i := 0; i < len(arg); i++ {
		if arg[0] != 45 && arg[i] < 48 || arg[i] > 57 {
			return false
		}
	}
	if arg[0] == 45 {
		negatif = true
		argSave = arg[1:]
	}
	for i := len(argSave) - 1; i >= 0; i-- {
		nb += uint(argSave[i]-48) * uint(puissance)
		puissance *= 10
	}
	if nb > 9223372036854775807 || nb > 9223372036854775808 && negatif == true {
		return false
	}
	return true
}

func validOperator(arg string) bool {
	if len(arg) == 1 {
		if arg == "+" || arg == "-" || arg == "*" || arg == "/" || arg == "%" {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) == 4 {
		os.Args = os.Args[1:]
		for i := 0; i < len(os.Args); i++ {
			if i != 1 {
				if validValue(os.Args[i]) == false {
					return
				}
			} else {
				if validOperator(os.Args[i]) == false {
					return
				}
			}
		}
		resultat := operation(os.Args[1], convertStringNumber(os.Args[0]), convertStringNumber(os.Args[2]))
		os.Stderr.WriteString(resultat)
		if resultat != "" {
			os.Stderr.WriteString("\n")
		}
	}
}
