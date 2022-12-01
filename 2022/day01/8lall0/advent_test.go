package main

import (
	"bytes"
	"io"
	"testing"
)

var input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test_part1(t *testing.T) {
	buf := bytes.NewReader([]byte(input))

	tests := []struct {
		name     string
		input    io.Reader
		expected Elf
	}{
		{
			name:  "test input",
			input: buf,
			expected: Elf{
				Num:      4,
				Calories: 24000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := part1(tt.input)
			if err != nil {
				t.Errorf("Unexpected error. Got: %v", err)
			}
			if n != tt.expected {
				t.Errorf("Expected: %v, got: %v", tt.expected, n)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	buf := bytes.NewReader([]byte(input))

	tests := []struct {
		name     string
		input    io.Reader
		expected []Elf
	}{
		{
			name:  "test input",
			input: buf,
			expected: []Elf{
				{
					Num:      4,
					Calories: 24000,
				},
				{
					Num:      3,
					Calories: 11000,
				},
				{
					Num:      5,
					Calories: 10000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Errorf("Unexpected error. Got: %v", err)
			}

			if len(got) != len(tt.expected) {
				t.Errorf("Number of elements, expected:%d, got: %d", len(tt.expected), len(got))
			}

			for i := 0; i < len(got); i++ {
				if got[i] != tt.expected[i] {
					t.Errorf("Elements at %d position, expected:%v, got: %v", i, tt.expected[i], got[i])
				}
			}
		})
	}
}
