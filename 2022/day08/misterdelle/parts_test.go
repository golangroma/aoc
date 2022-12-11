package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = `30373
25512
65332
33549
35390`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "21",
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
			input:    sample,
			expected: "8",
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

func TestIsVisibleEast(t *testing.T) {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	tt := []struct {
		name     string
		input    string
		row      int
		col      int
		expected bool
	}{
		//
		// 30373
		//
		{
			name:     "TestIsVisibleEast Row 0 Col 0",
			input:    sample,
			row:      0,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast Row 0 Col 1",
			input:    sample,
			row:      0,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 0 Col 2",
			input:    sample,
			row:      0,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 0 Col 3",
			input:    sample,
			row:      0,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast  Row 0 Col 4",
			input:    sample,
			row:      0,
			col:      4,
			expected: true,
		},
		//
		// 25512
		//
		{
			name:     "TestIsVisibleEast Row 1 Col 0",
			input:    sample,
			row:      1,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast Row 1 Col 1",
			input:    sample,
			row:      1,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 1 Col 2",
			input:    sample,
			row:      1,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast  Row 1 Col 3",
			input:    sample,
			row:      1,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 1 Col 4",
			input:    sample,
			row:      1,
			col:      4,
			expected: true,
		},
		//
		// 65332
		//
		{
			name:     "TestIsVisibleEast Row 2 Col 0",
			input:    sample,
			row:      2,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast Row 2 Col 1",
			input:    sample,
			row:      2,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast  Row 2 Col 2",
			input:    sample,
			row:      2,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 2 Col 3",
			input:    sample,
			row:      2,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast  Row 2 Col 4",
			input:    sample,
			row:      2,
			col:      4,
			expected: true,
		},
		//
		// 33549
		//

		{
			name:     "TestIsVisibleEast Row 3 Col 0",
			input:    sample,
			row:      3,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast Row 3 Col 1",
			input:    sample,
			row:      3,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 3 Col 2",
			input:    sample,
			row:      3,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 3 Col 3",
			input:    sample,
			row:      3,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 3 Col 4",
			input:    sample,
			row:      3,
			col:      4,
			expected: true,
		},
		//
		// 35390
		//

		{
			name:     "TestIsVisibleEast Row 4 Col 0",
			input:    sample,
			row:      4,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast Row 4 Col 1",
			input:    sample,
			row:      4,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 4 Col 2",
			input:    sample,
			row:      4,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleEast  Row 4 Col 3",
			input:    sample,
			row:      4,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleEast  Row 4 Col 4",
			input:    sample,
			row:      4,
			col:      4,
			expected: true,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			heightMap := loadHeightMap(input)

			if got := isVisibleEast(heightMap, tc.row, tc.col); got != tc.expected {
				t.Errorf("isVisibleEast() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsVisibleWest(t *testing.T) {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	tt := []struct {
		name     string
		input    string
		row      int
		col      int
		expected bool
	}{
		//
		// 30373
		//
		{
			name:     "TestIsVisibleWest Row 0 Col 0",
			input:    sample,
			row:      0,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest Row 0 Col 1",
			input:    sample,
			row:      0,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 0 Col 2",
			input:    sample,
			row:      0,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 0 Col 3",
			input:    sample,
			row:      0,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest  Row 0 Col 4",
			input:    sample,
			row:      0,
			col:      4,
			expected: false,
		},

		//
		// 25512
		//
		{
			name:     "TestIsVisibleWest Row 1 Col 0",
			input:    sample,
			row:      1,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest Row 1 Col 1",
			input:    sample,
			row:      1,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest  Row 1 Col 2",
			input:    sample,
			row:      1,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 1 Col 3",
			input:    sample,
			row:      1,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 1 Col 4",
			input:    sample,
			row:      1,
			col:      4,
			expected: false,
		},
		//
		// 65332
		//
		{
			name:     "TestIsVisibleWest Row 2 Col 0",
			input:    sample,
			row:      2,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest Row 2 Col 1",
			input:    sample,
			row:      2,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 2 Col 2",
			input:    sample,
			row:      2,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 2 Col 3",
			input:    sample,
			row:      2,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 2 Col 4",
			input:    sample,
			row:      2,
			col:      4,
			expected: false,
		},
		//
		// 33549
		//
		{
			name:     "TestIsVisibleWest Row 3 Col 0",
			input:    sample,
			row:      3,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest Row 3 Col 1",
			input:    sample,
			row:      3,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 3 Col 2",
			input:    sample,
			row:      3,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest  Row 3 Col 3",
			input:    sample,
			row:      3,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 3 Col 4",
			input:    sample,
			row:      3,
			col:      4,
			expected: true,
		},
		//
		// 35390
		//
		{
			name:     "TestIsVisibleWest Row 4 Col 0",
			input:    sample,
			row:      4,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest Row 4 Col 1",
			input:    sample,
			row:      4,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest  Row 4 Col 2",
			input:    sample,
			row:      4,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleWest  Row 4 Col 3",
			input:    sample,
			row:      4,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleWest  Row 4 Col 4",
			input:    sample,
			row:      4,
			col:      4,
			expected: false,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			heightMap := loadHeightMap(input)

			if got := isVisibleWest(heightMap, tc.row, tc.col); got != tc.expected {
				t.Errorf("isVisibleWest() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsVisibleNorth(t *testing.T) {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	tt := []struct {
		name     string
		input    string
		row      int
		col      int
		expected bool
	}{
		//
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleNorth Row 0 Col 0",
			input:    sample,
			row:      0,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth Row 0 Col 1",
			input:    sample,
			row:      0,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 0 Col 2",
			input:    sample,
			row:      0,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 0 Col 3",
			input:    sample,
			row:      0,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 0 Col 4",
			input:    sample,
			row:      0,
			col:      4,
			expected: true,
		},
		//
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleNorth Row 1 Col 0",
			input:    sample,
			row:      1,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth Row 1 Col 1",
			input:    sample,
			row:      1,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 1 Col 2",
			input:    sample,
			row:      1,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 1 Col 3",
			input:    sample,
			row:      1,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 1 Col 4",
			input:    sample,
			row:      1,
			col:      4,
			expected: false,
		},
		//
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleNorth Row 2 Col 0",
			input:    sample,
			row:      2,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth Row 2 Col 1",
			input:    sample,
			row:      2,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 2 Col 2",
			input:    sample,
			row:      2,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 2 Col 3",
			input:    sample,
			row:      2,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 2 Col 4",
			input:    sample,
			row:      2,
			col:      4,
			expected: false,
		},
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		{
			name:     "TestIsVisibleNorth Row 3 Col 0",
			input:    sample,
			row:      3,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth Row 3 Col 1",
			input:    sample,
			row:      3,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 3 Col 2",
			input:    sample,
			row:      3,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 3 Col 3",
			input:    sample,
			row:      3,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 3 Col 4",
			input:    sample,
			row:      3,
			col:      4,
			expected: true,
		},
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		{
			name:     "TestIsVisibleNorth Row 4 Col 0",
			input:    sample,
			row:      4,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth Row 4 Col 1",
			input:    sample,
			row:      4,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 4 Col 2",
			input:    sample,
			row:      4,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleNorth  Row 4 Col 3",
			input:    sample,
			row:      4,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleNorth  Row 4 Col 4",
			input:    sample,
			row:      4,
			col:      4,
			expected: false,
		},
	}
	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			heightMap := loadHeightMap(input)

			if got := isVisibleNorth(heightMap, tc.row, tc.col); got != tc.expected {
				t.Errorf("isVisibleNorth() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestIsVisibleSouth(t *testing.T) {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	tt := []struct {
		name     string
		input    string
		row      int
		col      int
		expected bool
	}{
		//
		// 30373
		// 25512
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleSouth Row 0 Col 0",
			input:    sample,
			row:      0,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth Row 0 Col 1",
			input:    sample,
			row:      0,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 0 Col 2",
			input:    sample,
			row:      0,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 0 Col 3",
			input:    sample,
			row:      0,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 0 Col 4",
			input:    sample,
			row:      0,
			col:      4,
			expected: false,
		},
		//
		// 25512
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleSouth Row 1 Col 0",
			input:    sample,
			row:      1,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth Row 1 Col 1",
			input:    sample,
			row:      1,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 1 Col 2",
			input:    sample,
			row:      1,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 1 Col 3",
			input:    sample,
			row:      1,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 1 Col 4",
			input:    sample,
			row:      1,
			col:      4,
			expected: false,
		},
		//
		// 65332
		// 33549
		// 35390
		//
		{
			name:     "TestIsVisibleSouth Row 2 Col 0",
			input:    sample,
			row:      2,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth Row 2 Col 1",
			input:    sample,
			row:      2,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 2 Col 2",
			input:    sample,
			row:      2,
			col:      2,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 2 Col 3",
			input:    sample,
			row:      2,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 2 Col 4",
			input:    sample,
			row:      2,
			col:      4,
			expected: false,
		},
		// 33549
		// 35390
		{
			name:     "TestIsVisibleSouth Row 3 Col 0",
			input:    sample,
			row:      3,
			col:      0,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth Row 3 Col 1",
			input:    sample,
			row:      3,
			col:      1,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 3 Col 2",
			input:    sample,
			row:      3,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth  Row 3 Col 3",
			input:    sample,
			row:      3,
			col:      3,
			expected: false,
		},
		{
			name:     "TestIsVisibleSouth  Row 3 Col 4",
			input:    sample,
			row:      3,
			col:      4,
			expected: true,
		},
		// 35390
		{
			name:     "TestIsVisibleSouth Row 4 Col 0",
			input:    sample,
			row:      4,
			col:      0,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth Row 4 Col 1",
			input:    sample,
			row:      4,
			col:      1,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth  Row 4 Col 2",
			input:    sample,
			row:      4,
			col:      2,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth  Row 4 Col 3",
			input:    sample,
			row:      4,
			col:      3,
			expected: true,
		},
		{
			name:     "TestIsVisibleSouth  Row 4 Col 4",
			input:    sample,
			row:      4,
			col:      4,
			expected: true,
		},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			heightMap := loadHeightMap(input)

			if got := isVisibleSouth(heightMap, tc.row, tc.col); got != tc.expected {
				t.Errorf("isVisibleSouth() = %v, want %v", got, tc.expected)
			}
		})
	}
}
