package main

func checkValue(val int, preamble []int) bool {
	for _, a := range preamble {
		for _, b := range preamble {
			if a != b && a+b == val {
				return true
			}
		}
	}
	return false
}

func findFirstXMASFailure(numbers []int, preambleLen int) int {
	for i := preambleLen; i < len(numbers); i++ {
		res := checkValue(numbers[i], numbers[i-preambleLen:i])
		if !res {
			return numbers[i]
		}
	}
	return -1
}
