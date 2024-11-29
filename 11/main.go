package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	IN_USE     = 1
	EMPTY      = -1
	NOT_A_SEAT = 0
)

func readData(file string) [][]int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	data := make([][]int, len(sData))

	for y, row := range sData {
		data[y] = make([]int, len(row))
		for x, char := range strings.Split(row, "") {
			if char == "L" {
				data[y][x] = EMPTY
			} else {
				data[y][x] = NOT_A_SEAT
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
