package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

var samples = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestPartOne(t *testing.T) {

	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example 1",
			input:    samples[0],
			expected: "7",
		}, {
			name:     "example 2",
			input:    samples[1],
			expected: "5",
		}, {
			name:     "example 3",
			input:    samples[2],
			expected: "6",
		}, {
			name:     "example 4",
			input:    samples[3],
			expected: "10",
		}, {
			name:     "example 5",
			input:    samples[4],
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
			input:    samples[0],
			expected: "19",
		}, {
			name:     "example 2",
			input:    samples[1],
			expected: "23",
		}, {
			name:     "example 3",
			input:    samples[2],
			expected: "23",
		}, {
			name:     "example 4",
			input:    samples[3],
			expected: "29",
		}, {
			name:     "example 5",
			input:    samples[4],
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
