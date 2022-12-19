package main

import (
	"strings"
	"testing"
)

// var sample string = `
// R 4
// U 4
// L 3
// D 1
// R 4
// D 1
// L 5
// R 2
// `

func TestPartOne(t *testing.T) {
	var sampleOne string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sampleOne,
			expected: "13",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//input := utils.SplitInput(tc.input)
			input := strings.Split(tc.input, "\n")
			if got := PartOne(input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	var sampleTwo string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sampleTwo,
			expected: "36",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.Split(tc.input, "\n")
			if got := PartTwo(input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
