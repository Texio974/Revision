package piscine

func IsMultiple(n int) bool {
	if n <= 0 {
		return false
	}
	if n%3 == 0 || n%7 == 0 {
		return true
	}
	return false
}
