package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := countThreeOneTraversal(data)

	const e = 7
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := mulTraversals(data)

	const e = 336
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
