package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := calcJolts(data)

	const e = 7 * 5
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
func TestA2(t *testing.T) {
	data := readData("test_data2")
	results := calcJolts(data)

	const e = 22 * 10
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := b(data)

	const e = 2
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
