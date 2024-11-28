package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	fmt.Println(data)
	results := calcOccupied(data)

	const e = 37
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
