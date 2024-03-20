package piscine

import (
	"fmt"
	"os"
)

func IsPrime(nb int) bool {
	if nb == 2 {
		return true
	}
	if nb%2 == 0 || nb <= 1 {
		return false
	}
	a := true
	for i := 3; i < nb; i++ {
		if nb%i == 0 {
			a = false
		}
	}
	return a
}

func addprimesum() {
	base := "0123456789"
	if len(os.Args) == 2 {
		arg := os.Args[1]
		val := 0
		for _, i := range arg {
			for j := 0; j < len(base); j++ {
				if i == rune(base[j]) {
					val += j
				}
				if i == '-' {
					fmt.Println(0)
					return
				}
			}
		}
		val2 := 0
		for i := 0; i < val+1; i++ {
			if IsPrime(i) {
				val2 += i
			}
		}
		fmt.Println(val2)
	} else {
		fmt.Println(0)
	}
}
