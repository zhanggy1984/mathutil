package mathutil

import "sort"

// Average returns the average of a slice of integers.
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// Max returns the maximum value in a slice of integers.
func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}

// Min returns the minimum value in a slice of integers.
func Min(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	sort.Ints(numbers)
	return numbers[0]
}
