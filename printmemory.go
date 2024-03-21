package piscine

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintHex(b byte) {
	if b/16 < 10 {
		z01.PrintRune(rune(b/16) + '0')
	} else {
		z01.PrintRune(rune(b/16) + 'a' - 10)
	}
	if b%16 < 10 {
		z01.PrintRune(rune(b%16) + '0')
	} else {
		z01.PrintRune(rune(b%16 + 'a' - 10))
	}
}

func PrintMemory(arr [10]byte) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if count == 3 {
			count = 0
			PrintHex(arr[i])
			z01.PrintRune('\n')
		} else if i == len(arr)-1 {
			PrintHex(arr[i])
			count++
		} else {
			PrintHex(arr[i])
			z01.PrintRune(' ')
			count++
		}
	}
	z01.PrintRune('\n')
	for _, i := range arr {
		if !unicode.IsGraphic(rune(i)) {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(i))
		}
	}
}
