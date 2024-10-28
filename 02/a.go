package main

import (
	"strings"
)

func validPws(pws []pwTest) int {
	validCounter := 0
	for _, pw := range pws {
		count := strings.Count(pw.pw, pw.char)
		if count >= pw.min && count <= pw.max {
			validCounter++
		}
	}
	return validCounter
}
