package piscine

func ItoaBase(n int, base string) string {
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return "NV"
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return "NV"
			}
		}
	}
	b := len(base)
	s := ""
	if n < 0 {
		n = -n
		s += "-"
	}
	if n == 0 {
		return "0"
	}
	for n != 0 {
		for i := 0; i < len(base); i++ {
			if n%b == i {
				s += string(base[i])
				n = (n - n%b) / b
			}
		}
	}
	return ReverseStr(s)
}
