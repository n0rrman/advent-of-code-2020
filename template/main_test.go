package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := a(data)

	const e = 1
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
