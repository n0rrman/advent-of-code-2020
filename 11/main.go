package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) [][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	data := make([][]int, len(sData))

	for y, row := range sData {
		fmt.Println(sData[y])
		data[y] = make([]int, len(row))
		for x, char := range strings.Split(row, "") {
			if char == "L" {
				data[y][x] = -1
			} else {
				data[y][x] = 0
			}
		}
	}

	return data
}

func main() {
	data := readData("data")

	// Part One
	results := calcOccupied(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = b(data)
	fmt.Println("Part two: ", results)
}
