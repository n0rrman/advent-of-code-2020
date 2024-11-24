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

	for i, d := range sData {
		num, _ := strconv.Atoi(d)
		data[i] = num
	}
	return data
}

func main() {
	data := readData("data")

	// Part One
	results := findFirstXMASFailure(data, 25)
	fmt.Println("Part one: ", results)

	// Part Two
	results = findEncryptionWeakness(data, 25)
	fmt.Println("Part two: ", results)
}
