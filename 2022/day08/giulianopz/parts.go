package main

import "fmt"

func PartOne(input []string) string {

	var visible int

	grid := make([][]int, len(input))
	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
	}

	width := len(grid)
	length := len(grid[0])
	visible += ((width + length) * 2) - 4

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {

			treeHeight := grid[row][col]

			ok := true
			up := row
			for up != 0 {
				if treeHeight <= grid[up-1][col] {
					ok = false
				}
				up--
			}
			if ok {
				visible++
				continue
			}

			ok = true
			down := row
			for down != len(grid)-1 {
				if treeHeight <= grid[down+1][col] {
					ok = false
				}
				down++
			}
			if ok {
				visible++
				continue
			}

			ok = true
			left := col
			for left != 0 {
				if treeHeight <= grid[row][left-1] {
					ok = false
				}
				left--
			}
			if ok {
				visible++
				continue
			}

			ok = true
			right := col
			for right != len(grid[row])-1 {
				if treeHeight <= grid[row][right+1] {
					ok = false
				}
				right++
			}
			if ok {
				visible++
				continue
			}
		}
	}

	return fmt.Sprintf("%d", visible)
}

func PartTwo(input []string) string {
	return ""
}
