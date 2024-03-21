package piscine

func AppendRange(min, max int) []int {
	res := []int{}
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
