package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := sumUniqueVotes(data)

	const e = 11
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := sumUnanimousVotes(data)

	const e = 6
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
