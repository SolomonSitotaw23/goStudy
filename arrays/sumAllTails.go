package arrays

func SumAllTails(numbersToSum ...[]int) []int {

	var sumOfTails []int

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sumOfTails = append(sumOfTails, 0)
		} else {
			tail := number[1:]
			sumOfTails = append(sumOfTails, Sum(tail))
		}
	}

	return sumOfTails
}
