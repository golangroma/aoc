package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = `
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "13",
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

func TestParseInput(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple ints",
			input:    "[1,1,3,1,1]",
			expected: "[1,1,3,1,1]",
		},
		{
			name:     "simple nested",
			input:    "[[1],[2,3,4]]",
			expected: "[[1],[2,3,4]]",
		},
		{
			name:     "complex nested",
			input:    "[[4,4],4,4]",
			expected: "[[4,4],4,4]",
		},
		{
			name:     "complex nested 2",
			input:    "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			expected: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
		},
		{
			name:     "complex nested 3",
			input:    "[[[]]]",
			expected: "[[[]]]",
		},
		{
			name:     "complex nested 4",
			input:    "[[]]",
			expected: "[[]]",
		},
		{
			name:     "complex nested 5",
			input:    "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			expected: "[1,[2,[3,[4,[5,6,0]]]],8,9]",
		},
		{
			name:     "complex nested 6",
			input:    "[]",
			expected: "[]",
		},
		{
			name:     "complex nested 7",
			input:    "[3]",
			expected: "[3]",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.input
			got := Node{}
			ParseInput(input, 1, &got)

			str := fmt.Sprintf("[%s]", got.String())

			if !reflect.DeepEqual(str, tc.expected) {
				t.Errorf("ParseInput() = %v, want %v", str, tc.expected)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tt := []struct {
		name     string
		inputA   string
		inputB   string
		expected bool
	}{
		{
			name:     "test one",
			inputA:   "[1,1,3,1,1]",
			inputB:   "[1,1,5,1,1]",
			expected: true,
		},
		{
			name:     "test two",
			inputA:   "[[1],[2,3,4]]",
			inputB:   "[[1],4]",
			expected: true,
		},
		{
			name:     "test three",
			inputA:   "[9]",
			inputB:   "[[8,7,6]]",
			expected: false,
		},
		{
			name:     "test four",
			inputA:   "[[4,4],4,4]",
			inputB:   "[[4,4],4,4,4]",
			expected: true,
		},
		{
			name:     "test five",
			inputA:   "[7,7,7,7]",
			inputB:   "[7,7,7]",
			expected: false,
		},
		{
			name:     "test six",
			inputA:   "[]",
			inputB:   "[3]",
			expected: true,
		},
		{
			name:     "test seven",
			inputA:   "[[[]]]",
			inputB:   "[[]]",
			expected: false,
		},
		{
			name:     "test eight",
			inputA:   "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			inputB:   "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a := Node{}
			ParseInput(tc.inputA, 1, &a)
			b := Node{}
			ParseInput(tc.inputB, 1, &b)
			if got, _ := Compare(&a, &b); got != tc.expected {
				t.Errorf("Compare() = %v, want %v", got, tc.expected)
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
			expected: "140",
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
