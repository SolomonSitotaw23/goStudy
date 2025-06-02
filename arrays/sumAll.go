package arrays

func SumAll(numbersToSum ...[]int) (sums []int) {

	sums = make([]int, len(numbersToSum))

	for i, number := range numbersToSum {
		sums[i] = Sum(number)
	}

	return
}
