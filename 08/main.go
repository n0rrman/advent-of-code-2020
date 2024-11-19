package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	program := make([]instruction, len(sData))
	for iter, p := range sData {
		split := strings.Split(p, " ")
		t := split[0]
		v, _ := strconv.Atoi(split[1])
		i := instruction{
			InstrType: t,
			Value:     v,
			Executed:  false,
		}
		program[iter] = i
	}

	return program
}

func main() {
	program := readData("data")

	// Part One
	results := executeProgram(program)
	fmt.Println("Part one: ", results)

	// Part Two
	results = b(program)
	fmt.Println("Part two: ", results)
}
