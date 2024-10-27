package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	fmt.Println(data)

	// Part One
	results := 1

	const e = 2
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	results := 0

	const e = 2
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
