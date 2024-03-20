package piscine

func ReverseStr(s string) string {
	rev := ""
	if s[0] == '-' {
		rev += "-"
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return rev
}
func Itoa(n int) string {
	base := "0123456789"
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
			if n%10 == i {
				s += string(base[i])
				n = (n - n%10) / 10
			}
		}
	}
	return ReverseStr(s)
}
