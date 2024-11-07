package main

import (
	"math"
)

func findHighestSeatID(directions [][]string, rows, seats int) int {
	maxId := 0
	for _, ticket := range directions {
		minRow, maxRow := 0, rows-1
		minSeat, maxSeat := 0, seats-1

		for _, dir := range ticket {
			switch dir {
			case "F":
				maxRow = int(math.Ceil(float64(maxRow+minRow)/2)) - 1
			case "B":
				minRow = int(math.Ceil((float64(maxRow + minRow)) / 2))
			case "L":
				maxSeat = int(math.Ceil(float64(maxSeat+minSeat)/2)) - 1
			case "R":
				minSeat = int(math.Ceil(float64(maxSeat+minSeat) / 2))
			}
		}
		id := minRow*8 + minSeat
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}
