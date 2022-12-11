package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = ``

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example 1",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: "5",
		},
		{
			name:     "example 2",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: "6",
		},
		{
			name:     "example 3",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: "10",
		},
		{
			name:     "example 4",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: "11",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
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
			name:     "example 1",
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: "19",
		},
		{
			name:     "example 2",
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: "23",
		},
		{
			name:     "example 3",
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: "23",
		},
		{
			name:     "example 4",
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: "29",
		},
		{
			name:     "example 5",
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: "26",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			if got := PartTwo(input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
