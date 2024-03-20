package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for !IsPrime(nb) {
		nb--
	}
	return nb
}
