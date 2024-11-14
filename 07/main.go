package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type bag struct {
	colour string
	bags   []string
}

func readData(file string) map[string]bag {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Replace(string(body[:]), ".", "", -1)
	sData = strings.Replace(sData, " bags", "", -1)
	sData = strings.Replace(sData, " bag", "", -1)
	dataString := strings.Split(sData, "\n")

	bagList := make(map[string]bag)
	for _, bagString := range dataString {

		// Colour
		colourSplit := strings.Split(bagString, " contain ")
		colour := colourSplit[0]
		bag := bag{colour: colour, bags: []string{}}

		// Contains
		if colourSplit[1] == "no other" {
			bagList[colour] = bag
			continue
		}
		containsSplit := strings.Split(colourSplit[1], ", ")
		for _, containedBag := range containsSplit {
			amount := int(containedBag[0] - 48)
			colour := containedBag[2:]

			for i := 0; i < amount; i++ {
				bag.bags = append(bag.bags, colour)
			}
		}
		bagList[colour] = bag
	}

	return bagList
}

func main() {
	data := readData("data")

	// Part One
	results := countBagColours("shiny gold", data)
	fmt.Println("Part one: ", results)

	// Part Two
	results = 2
	fmt.Println("Part two: ", results)
}
