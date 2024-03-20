package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	for _, i := range ItoaBase(nbr, base) {
		z01.PrintRune(i)
	}
}
