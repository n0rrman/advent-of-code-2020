package main

import (
	"math"
)

func findSeatID(directions [][]string, rows, seats int) int {
	seatList := make([]bool, rows*seats)
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
		seatList[minRow*8+minSeat] = true
	}
	return findSeat(seatList)
}

func findSeat(seats []bool) int {
	for i := range seats {
		if i < 2 {
			continue
		}

		if seats[i-2] && !seats[i-1] && seats[i] {
			return i - 1
		}
	}
	return -1
}
