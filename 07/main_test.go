package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := countBagColours("shiny gold", data)

	const e = 4
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB1(t *testing.T) {
	data := readData("test_data")
	results := countContainedBags("shiny gold", data)

	const e = 32
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB2(t *testing.T) {
	data := readData("test_data2")
	results := countContainedBags("shiny gold", data)

	const e = 126
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
