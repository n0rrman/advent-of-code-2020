package main

func countThreeOneTraversal(g grid) int {
	count := 0
	for i := 0; i < len(g); i++ {
		y := i
		x := (i * 3) % len(g[i])
		if g[y][x] {
			count++
		}
	}
	return count
}
