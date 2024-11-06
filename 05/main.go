package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]bool {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	dirArray := make([][]bool, len(sData))

	for y, row := range sData {
		dirArray[y] = make([]bool, len(row))
		for x, val := range strings.Split(row, "") {
			boolVal := val == "F" || val == "R"
			dirArray[y][x] = boolVal
		}
	}

	return dirArray
}

func main() {
	data := readData("data")
	_ = data

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
