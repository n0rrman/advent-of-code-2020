package main

func sumMinMax(numbers []int) int {
	min := numbers[0]
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return min + max
}

func findEncryptionWeakness(numbers []int, preambleLen int) int {
	failedVal := findFirstXMASFailure(numbers, preambleLen)
	lowerIndex := 0
	upperIndex := 1
	compSum := numbers[lowerIndex] + numbers[upperIndex]

	for compSum != failedVal {
		if compSum > failedVal {
			compSum -= numbers[lowerIndex]
			lowerIndex++
		} else {
			upperIndex++
			compSum += numbers[upperIndex]
		}
	}
	return sumMinMax(numbers[lowerIndex : upperIndex+1])

}
