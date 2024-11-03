package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type passportData struct {
	byr string //(Birth Year)
	iyr string //(Issue Year)
	eyr string //(Expiration Year)
	hgt string //(Height)
	hcl string //(Hair Color)
	ecl string //(Eye Color)
	pid string //(Passport ID)
	cid string //(Country ID)
}

func readData(file string) []passportData {
	body, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	sData := strings.Split(string(body[:]), "\n\n")

	pd := make([]passportData, len(sData))
	for i, s := range sData {
		data := strings.Split(strings.Replace(s, "\n", " ", -1), " ")
		passportData := passportData{}
		for _, dataType := range data {
			dt := strings.Split(dataType, ":")
			switch dt[0] {
			case "byr":
				passportData.byr = dt[1]
			case "iyr":
				passportData.iyr = dt[1]
			case "eyr":
				passportData.eyr = dt[1]
			case "hgt":
				passportData.hgt = dt[1]
			case "hcl":
				passportData.hcl = dt[1]
			case "ecl":
				passportData.ecl = dt[1]
			case "pid":
				passportData.pid = dt[1]
			case "cid":
				passportData.cid = dt[1]
			}
		}
		pd[i] = passportData
	}
	return pd
}

func main() {
	data := readData("test_data")
	fmt.Println(data)
	_ = data

	// Part One
	results := "part one"
	fmt.Println("Part one: ", results)

	// Part Two
	results = "part two"
	fmt.Println("Part two: ", results)
}
