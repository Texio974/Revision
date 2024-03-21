package piscine

func SumArray(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
