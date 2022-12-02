package main

import (
	"testing"
)

var sample []string = []string{"1", "2", "3", "", "2", "", "3", "", "11", ""}

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "11",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartOne(tc.input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "20",
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
