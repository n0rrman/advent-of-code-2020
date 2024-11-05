package main

import (
	"strconv"
)

func validateNum(y string, min, max int) bool {
	if y == "" {
		return false
	}
	numYear, err := strconv.Atoi(y)
	if err != nil {
		return false
	}

	if numYear < min || numYear > max {
		return false
	}

	return true
}

func validateHeight(h string) bool {
	if h == "" {
		return false
	}

	cmOrIn := h[len(h)-2:]
	if cmOrIn == "cm" {
		return validateNum(h[:len(h)-2], 150, 193)
	} else if cmOrIn == "in" {
		return validateNum(h[:len(h)-2], 59, 76)
	}

	return false
}

func validateHex(h string) bool {
	if h == "" {
		return false
	}
	if h[0] != '#' || len(h) != 7 {
		return false
	}

	for _, r := range h[1:] {
		if r < 48 || r > 122 {
			return false
		}
	}

	return true
}

func validateEcl(e string) bool {
	if e == "" {
		return false
	}
	return e == "amb" || e == "blu" || e == "brn" || e == "gry" || e == "grn" || e == "hzl" || e == "oth"
}

func validatePid(p string) bool {
	if p == "" {
		return false
	}
	if len(p) != 9 {
		return false
	}
	for _, r := range p {
		if r < 48 || r > 57 {
			return false
		}
	}
	return true
}

func countValidatedIds(pd []passportData) int {
	count := 0
	for _, id := range pd {
		if validateNum(id.byr, 1920, 2002) &&
			validateNum(id.iyr, 2010, 2020) &&
			validateNum(id.eyr, 2020, 2030) &&
			validateHeight(id.hgt) &&
			validateHex(id.hcl) &&
			validateEcl(id.ecl) &&
			validatePid(id.pid) {
			count++
		}
	}
	return count
}
