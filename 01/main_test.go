package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := findReportVal(data)

	const e = 514579
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := findSecondReportVal(data)

	const e = 241861950
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
