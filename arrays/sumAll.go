package arrays

// accepts any number of slices and return a
// slice with the sum of each slice
func SumAll(numbersToSum ...[]int) (sums []int) {

	sums = make([]int, len(numbersToSum))

	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return
}
