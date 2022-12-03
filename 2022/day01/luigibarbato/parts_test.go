package main

import (
	"testing"
)

var sample string = ``

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "Successfull test",
			input: []string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				"",
			},
			expected: 24000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := PartOne(tc.input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			} else if err != nil {
				t.Errorf("PartOne() error = %v", err)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "Successfull test",
			input: []string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				"",
			},
			expected: 45000,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := PartTwo(tc.input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			} else if err != nil {
				t.Errorf("PartTwo() error = %v", err)
			}
		})
	}
}
