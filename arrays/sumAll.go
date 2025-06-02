package arrays

// accepts any number of slices and return a
// slice with the sum of each slice

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return
}
