package main

import (
	"testing"
)

func TestA(t *testing.T) {
	program := readData("test_data")
	results := executeProgram(program)

	const e = 5
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := executeModifiedPrograms(data)

	const e = 8
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
