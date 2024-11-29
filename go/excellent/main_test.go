package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2, "Even"},
		{3, "Odd"},
		{0, "Even"},
		{-1, "Odd"},
		{-2, "Odd"},
	}

	for _, test := range tests {
		result := EvenOrOdd(test.input)
		if result != test.expected {
			t.Errorf("For input %d, expected %s but got %s", test.input, test.expected, result)
		}
	}
}