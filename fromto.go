package piscine

func FromTo(from int, to int) string {
	if from >= 100 || to >= 100 || from < 0 || to < 0 {
		return "Invalid\n"
	}
	s := ""
	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 {
				s += "0"
				s += string(rune(i + 48))
			} else {
				s += string(rune(i/10 + 48))
				s += string(rune(i%10 + 48))
			}
			if i != to {
				s += ", "
			}
		}
		s += "\n"
		return s
	}
	for i := from; i <= to; i++ {
		if i < 10 {
			s += "0"
			s += string(rune(i + 48))
		} else {
			s += string(rune(i/10 + 48))
			s += string(rune(i%10 + 48))
		}
		if i != to {
			s += ", "
		}
	}
	s += "\n"
	return s
}
