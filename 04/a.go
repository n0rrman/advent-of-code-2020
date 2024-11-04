package main

func countValidIds(pd []passportData) int {
	count := 0
	for _, id := range pd {
		if id.byr != "" &&
			id.iyr != "" &&
			id.eyr != "" &&
			id.hgt != "" &&
			id.hcl != "" &&
			id.ecl != "" &&
			id.pid != "" {
			count++
		}
	}
	return count
}
