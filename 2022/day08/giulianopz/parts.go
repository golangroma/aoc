package main

import (
	"fmt"
)

func PartOne(input []string) string {

	grid := makeGrid(input)

	width := len(grid)
	length := len(grid[0])
	visibleTrees := ((width + length) * 2) - 4

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {

			treeHeight := grid[row][col]

			visible := true
			up := row
			for up != 0 {
				if treeHeight <= grid[up-1][col] {
					visible = false
				}
				up--
			}
			if visible {
				visibleTrees++
				continue
			}

			visible = true
			down := row
			for down != len(grid)-1 {
				if treeHeight <= grid[down+1][col] {
					visible = false
				}
				down++
			}
			if visible {
				visibleTrees++
				continue
			}

			visible = true
			left := col
			for left != 0 {
				if treeHeight <= grid[row][left-1] {
					visible = false
				}
				left--
			}
			if visible {
				visibleTrees++
				continue
			}

			visible = true
			right := col
			for right != len(grid[row])-1 {
				if treeHeight <= grid[row][right+1] {
					visible = false
				}
				right++
			}
			if visible {
				visibleTrees++
				continue
			}
		}
	}

	return fmt.Sprintf("%d", visibleTrees)
}

func PartTwo(input []string) string {

	grid := makeGrid(input)

	var maxScenicScore int

	for row := 1; row < len(grid)-1; row++ {

		for col := 1; col < len(grid[row])-1; col++ {

			treeHeight := grid[row][col]
			var scenicScore int

			var trees int
			up := row
			for up != 0 {
				trees++
				if treeHeight <= grid[up-1][col] {
					break
				}
				up--
			}
			trees, scenicScore = score(trees, scenicScore)

			down := row
			for down != len(grid)-1 {
				trees++
				if treeHeight <= grid[down+1][col] {
					break
				}
				down++
			}
			trees, scenicScore = score(trees, scenicScore)

			left := col
			for left != 0 {
				trees++
				if treeHeight <= grid[row][left-1] {
					break
				}
				left--
			}
			trees, scenicScore = score(trees, scenicScore)

			right := col
			for right != len(grid[row])-1 {
				trees++
				if treeHeight <= grid[row][right+1] {
					break
				}
				right++
			}
			trees, scenicScore = score(trees, scenicScore)

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return fmt.Sprintf("%d", maxScenicScore)
}

func makeGrid(input []string) [][]int {
	grid := make([][]int, len(input))
	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
	}
	return grid
}

func score(trees int, scenicScore int) (int, int) {
	if trees != 0 {
		if scenicScore == 0 {
			scenicScore = 1
		}
		return 0, trees * scenicScore
	}
	return 0, scenicScore
}
