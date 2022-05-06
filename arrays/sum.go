package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			// The sum of an empty slice's tail is 0
			sums = append(sums, 0)
			continue
		}
		sums = append(sums, Sum(numbers[1:]))
	}
	return sums
}
