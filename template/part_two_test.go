package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example",
			input:    []string{},
			expected: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartTwo(tc.input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
