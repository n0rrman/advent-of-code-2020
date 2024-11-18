package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type instruction struct {
	InstrType string
	Value     int
	Executed  bool
}

func readData(file string) []instruction {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n")

	return sData
}

func main() {
	data := readData("data")

	// Part One
	results := a(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = b(data)
	fmt.Println("Part two: ", results)
}
