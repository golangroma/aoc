package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const (
	head  = "H"
	tail  = "T"
	up    = "U"
	down  = "D"
	right = "R"
	left  = "L"
)

func PartOne(input []string) string {

	visited := make(map[string]int, 0)
	grid := make([][]string, 1000)
	for i := range grid {
		grid[i] = make([]string, 1000)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	cmdRgx, _ := regexp.Compile("(\\w)\\s(\\d+)")
	start := 501
	grid[start][start] = "s"
	hRow := start
	hCol := start
	tRow := hRow
	tCol := hCol
	visited[fmt.Sprintf("%d%d", tRow, tCol)] += 1
	for _, line := range input {

		if cmdRgx.MatchString(line) {

			matches := cmdRgx.FindAllStringSubmatch(line, -1)
			direction := matches[0][1]
			steps, err := strconv.Atoi(matches[0][2])
			if err != nil {
				log.Fatal(err)
			}

			for steps != 0 {

				grid[hRow][hCol] = "."
				switch direction {
				case right:
					hCol += 1
				case left:
					hCol -= 1
				case up:
					hRow -= 1
				case down:
					hRow += 1
				}
				grid[hRow][hCol] = head

				if !adjacent(tRow, tCol, hRow, hCol) {

					grid[tRow][tCol] = "#"
					if diagonally(tRow, tCol, hRow, hCol) {
						//move diagonally towards H

						//if H above T
						if hRow < tRow {
							tRow -= 1
						} else {
							//else H below T
							tRow += 1
						}
						//if H to the left of T
						if hCol < tCol {
							tCol -= 1
						} else {
							tCol += 1
						}
					} else {
						//move one step towards H
						switch direction {
						case right:
							tCol += 1
						case left:
							tCol -= 1
						case up:
							tRow -= 1
						case down:
							tRow += 1
						}
					}
					grid[tRow][tCol] = tail
					visited[fmt.Sprintf("%d%d", tRow, tCol)] += 1
				}

				steps--
			}
		}
	}
	return fmt.Sprintf("%d", len(visited))
}

func adjacent(tRow, tCol, hRow, hCol int) bool {
	if tRow == hRow && tCol == hCol {
		return true
	}
	if math.Abs(float64(tRow-hRow)) >= 2 || math.Abs(float64(tCol-hCol)) >= 2 {
		return false
	}
	return true
}

func diagonally(tRow, tCol, hRow, hCol int) bool {
	return tRow != hRow && tCol != hCol
}

func print(grid [][]string) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}

func PartTwo(input []string) string {
	return ""
}
