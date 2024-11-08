package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := findHighestSeatID(data, 128, 8)

	const e = 820
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
