package tddexample

func Sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}

	return s
}

// returning a new slice containing the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	result := make([]int, len(numbersToSum))

	for i, ints := range numbersToSum {
		result[i] = Sum(ints)
	}

	return result

}
