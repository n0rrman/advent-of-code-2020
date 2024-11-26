package main

import (
	"sort"
)

func calcJolts(adapters []int) int {
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	diffCounts := make([]int, 4)
	diffCounts[3] = 1
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		diffCounts[diff]++
		if diff > 3 {
			return -1
		}
	}
	return diffCounts[1] * diffCounts[3]
}
