package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 32 {
			if s[i+1] >= 'a' && s[i+1] <= 'z' {
				return false
			}
		}
	}
	return true
}
