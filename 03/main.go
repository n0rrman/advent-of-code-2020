package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type grid [][]bool

func readData(file string) grid {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	g := make(grid, len(sData))
	for y, row := range sData {
		sRow := strings.Split(row, "")
		g[y] = make([]bool, len(sRow))
		for x, entry := range sRow {
			g[y][x] = entry == "#"
		}
	}

	return g
}

func main() {
	data := readData("data")

	// Part One
	results := countThreeOneTraversal(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = mulTraversals(data)
	fmt.Println("Part two: ", results)
}
