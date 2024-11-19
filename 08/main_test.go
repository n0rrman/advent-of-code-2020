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
	results := b(data)

	const e = 2
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
