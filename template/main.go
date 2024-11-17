package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(file string) []string {
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
