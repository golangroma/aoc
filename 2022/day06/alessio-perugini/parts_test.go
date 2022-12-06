package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    `bvwbjplbgvbhsrlpgdmjqwftvncz`,
			expected: "5",
		},
		{
			name:     "2",
			input:    `nppdvjthqldpwncqszvftbrmjlhg`,
			expected: "6",
		},
		{
			name:     "2",
			input:    `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
			expected: "10",
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
			name:     "example",
			input:    `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
			expected: "19",
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
