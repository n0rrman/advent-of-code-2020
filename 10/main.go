package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readData(file string) []int {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")
	data := make([]int, len(sData))
	for i, s := range sData {
		val, _ := strconv.Atoi(s)
		data[i] = val
	}
	return data
}

func main() {
	data := readData("data")

	// Part One
	results := calcJolts(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = b(data)
	fmt.Println("Part two: ", results)
}
