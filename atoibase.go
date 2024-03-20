package piscine

func AtoiBase(s string, base string) int {
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	b := len(base)
	val := 0
	for _, i := range s {
		for j := 0; j < len(base); j++ {
			if i == rune(base[j]) {
				val = val*b + j
			}
		}
	}
	return val
}
