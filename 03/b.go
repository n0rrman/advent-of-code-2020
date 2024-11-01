package main

func countTraversal(g grid, xShift, yShift int) int {
	count := 0
	for i := 0; i < len(g); i += yShift {
		y := i
		x := ((i / yShift) * xShift) % len(g[i])
		if g[y][x] {
			count++
		}
	}
	return count
}

func mulTraversals(g grid) int {
	shifts := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	product := 1
	for _, shift := range shifts {
		product *= countTraversal(g, shift[0], shift[1])
	}
	return product
}
