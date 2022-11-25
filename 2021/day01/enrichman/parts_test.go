package main

import (
	"testing"

	"github.com/golangroma/aoc/util"
)

var sample string = `
199
200
208
210
200
207
240
269
260
263
`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "7",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := util.SplitInput(tc.input)
			if got := PartOne(input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "5",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := util.SplitInput(tc.input)
			if got := PartTwo(input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
