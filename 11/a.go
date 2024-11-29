package main

func countAdjacent(s [][]int, y, x int) int {
	count := 0
	if x > 0 && s[y][x-1] == IN_USE {
		count++
	}
	if x < len(s[0])-1 && s[y][x+1] == IN_USE {
		count++
	}
	if y > 0 {
		if s[y-1][x] == IN_USE {
			count++
		}
		if x > 0 && s[y-1][x-1] == IN_USE {
			count++
		}
		if x < len(s[0])-1 && s[y-1][x+1] == IN_USE {
			count++
		}
	}
	if y < len(s)-1 {
		if s[y+1][x] == IN_USE {
			count++
		}
		if x > 0 && s[y+1][x-1] == IN_USE {
			count++
		}
		if x < len(s[0])-1 && s[y+1][x+1] == IN_USE {
			count++
		}
	}
	return count
}

func oneStep(s [][]int) ([][]int, bool) {
	changed := false
	newSeats := make([][]int, len(s))
	for y := range s {
		newSeats[y] = make([]int, len(s[y]))
		for x := range s[y] {
			numAdj := countAdjacent(s, y, x)
			seat := s[y][x]
			if seat == IN_USE && numAdj >= 4 {
				newSeats[y][x] = EMPTY
				changed = true
			} else if seat == EMPTY && numAdj == 0 {
				newSeats[y][x] = IN_USE
				changed = true
			} else {
				newSeats[y][x] = seat
			}
		}
	}
	return newSeats, changed
}

func stabiliseSeats(s [][]int) [][]int {
	seats := s
	changed := true
	for changed {
		seats, changed = oneStep(seats)
	}
	return seats
}

func countSeats(s [][]int) int {
	count := 0
	for y := range s {
		for x := range s[y] {
			if s[y][x] == IN_USE {
				count++
			}
		}
	}
	return count
}

func calcOccupied(seats [][]int) int {
	seats = stabiliseSeats(seats)
	count := countSeats(seats)

	return count
}
