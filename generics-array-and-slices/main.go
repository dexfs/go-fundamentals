package generics_as

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

// Sum calculates the total from slice of numbers
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbers, sumTail, []int{})
}
