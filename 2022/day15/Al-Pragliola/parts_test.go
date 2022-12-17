package main

import (
	"strconv"
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		y        int
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			y:        10,
			expected: "26",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)

			pairs := GetPairs(input)
			set := make(map[string]bool)

			for i := 0; i < len(pairs); i++ {
				GetAllPointsFromAToBAtY(
					int(real(pairs[i].Sensor)),
					int(imag(pairs[i].Sensor)),
					int(real(pairs[i].Beacon)),
					int(imag(pairs[i].Beacon)),
					tc.y,
					&set,
				)
			}

			got := strconv.Itoa(len(set))

			if got != tc.expected {
				t.Errorf("GetNumPointsAtY() = %v, want %v", got, tc.expected)
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
			expected: "56000011",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)

			got := ""

			max := 20

			pairs := GetPairs(input)

			for i := 0; i <= max; i++ {
				ranges := FillRanges(pairs, i, max)

				gap := FindGapInRanges(ranges, max)

				if gap != nil {
					got = strconv.Itoa(gap[0]*4000000 + i)
					break
				}
			}

			if got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
