package array

func Sum(numbers []int) int {
	//sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}
	//return sum

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
