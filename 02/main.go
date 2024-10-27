package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pwTest struct {
	min  int
	max  int
	char string
	pw   string
}

func readData(file string) []pwTest {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	tests := make([]pwTest, len(sData))
	for i, test := range sData {
		parts := strings.Split(test, " ")

		tests[i].char = strings.Replace(parts[1], ":", "", 1)
		tests[i].pw = parts[2]

		minMax := strings.Split(parts[0], "-")
		tests[i].min, _ = strconv.Atoi(minMax[0])
		tests[i].max, _ = strconv.Atoi(minMax[1])
	}

	return tests
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
