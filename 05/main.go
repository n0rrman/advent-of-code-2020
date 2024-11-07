package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]string {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	dirArray := make([][]string, len(sData))

	for y, row := range sData {
		dirArray[y] = make([]string, len(row))
		for x, val := range strings.Split(row, "") {
			dirArray[y][x] = val
		}
	}

	return dirArray
}

func main() {
	data := readData("data")

	// Part One
	results := findHighestSeatID(data, 128, 8)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 2
	fmt.Println("Part two: ", results)
}
