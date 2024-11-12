package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type group []person
type person []vote
type vote string

func readData(file string) []group {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n\n")

	groups := make([]group, len(sData))
	for z, g := range sData {
		persons := strings.Split(g, "\n")
		groups[z] = make([]person, len(persons))
		for y, p := range persons {
			votes := strings.Split(p, "")
			groups[z][y] = make([]vote, len(votes))
			for x, v := range votes {
				groups[z][y][x] = vote(v)
			}
		}
	}
	return groups
}

func main() {
	data := readData("data")

	// Part One
	results := sumUniqueVotes(data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = sumUnanimousVotes(data)
	fmt.Println("Part two: ", results)
}
