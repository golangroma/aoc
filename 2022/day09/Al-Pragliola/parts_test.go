package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sampleOne string = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var sampleTwo string = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

var sampleThree string = `
U 17
R 2
D 3
L 5
D 3
R 5
D 4
L 5
U 4
R 5
U 3
L 2`

func TestPartOne(t *testing.T) {
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
			input := utils.SplitInput(tc.input)
			if got := PartOne(input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestFollow(t *testing.T) {
	tt := []struct {
		name         string
		posH         Pos
		posT         Pos
		expectedPosT Pos
	}{
		{
			name:         "right",
			posH:         Pos{1, 3},
			posT:         Pos{1, 1},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "left",
			posH:         Pos{1, 1},
			posT:         Pos{1, 3},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "up",
			posH:         Pos{1, 1},
			posT:         Pos{3, 1},
			expectedPosT: Pos{2, 1},
		},
		{
			name:         "down",
			posH:         Pos{3, 1},
			posT:         Pos{1, 1},
			expectedPosT: Pos{2, 1},
		},
		{
			name:         "diag top",
			posH:         Pos{1, 2},
			posT:         Pos{3, 1},
			expectedPosT: Pos{2, 2},
		},
		{
			name:         "diag right",
			posH:         Pos{2, 3},
			posT:         Pos{3, 1},
			expectedPosT: Pos{2, 2},
		},
		{
			name:         "diag no move - 1",
			posH:         Pos{0, 1},
			posT:         Pos{1, 2},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "diag no move - 2",
			posH:         Pos{0, 3},
			posT:         Pos{1, 2},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "diag no move - 3",
			posH:         Pos{2, 3},
			posT:         Pos{1, 2},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "diag no move - 4",
			posH:         Pos{2, 1},
			posT:         Pos{1, 2},
			expectedPosT: Pos{1, 2},
		},
		{
			name:         "diag move - 1",
			posH:         Pos{2, 0},
			posT:         Pos{1, 2},
			expectedPosT: Pos{2, 1},
		},
		{
			name:         "diag move - 2",
			posH:         Pos{3, 1},
			posT:         Pos{1, 2},
			expectedPosT: Pos{2, 1},
		},
		{
			name:         "diag move - 3",
			posH:         Pos{3, 3},
			posT:         Pos{2, 1},
			expectedPosT: Pos{3, 2},
		},
		{
			name:         "diag move - 4",
			posH:         Pos{3, 3},
			posT:         Pos{1, 2},
			expectedPosT: Pos{2, 3},
		},
		{
			name:         "diag move - 5",
			posH:         Pos{3, 3},
			posT:         Pos{1, 1},
			expectedPosT: Pos{2, 2},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			grid := make([][]int, 10)
			for i := range grid {
				grid[i] = make([]int, 10)
			}

			Follow(tc.posH, &tc.posT, &grid, true)

			if got := tc.posT; got != tc.expectedPosT {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expectedPosT)
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
			input:    sampleOne,
			expected: "1",
		},
		{
			name:     "example",
			input:    sampleTwo,
			expected: "36",
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
