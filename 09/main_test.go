package main

import (
	"testing"
)

func TestA(t *testing.T) {
	data := readData("test_data")
	results := findFirstXMASFailure(data, 5)

	const e = 127
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}

func TestB(t *testing.T) {
	data := readData("test_data")
	results := findEncryptionWeakness(data, 5)

	const e = 62
	if results != e {
		t.Errorf("Expected %v, but got %v", e, results)
	}
}
