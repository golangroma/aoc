package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const (
	/* knots */
	head  = "H"
	tail  = "T"
	one   = "1"
	two   = "2"
	three = "3"
	four  = "4"
	five  = "5"
	six   = "6"
	seven = "7"
	eigth = "8"
	/* directions */
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

	start := 501
	grid[start][start] = "s"

	hRow := start
	hCol := start
	tRow := hRow
	tCol := hCol
	visited[fmt.Sprintf("%d%d", tRow, tCol)] += 1

	cmdRgx, _ := regexp.Compile("(\\w)\\s(\\d+)")
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

func PartTwo(input []string) string {
	visited := make(map[string]int, 0)

	grid := make([][]string, 1000)
	for i := range grid {
		grid[i] = make([]string, 1000)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	start := 501
	grid[start][start] = "s"

	headR := start
	headC := start
	tailR := headR
	tailC := headC
	visited[fmt.Sprintf("%d%d", tailR, tailC)] += 1

	oneR, twoR, threeR, fourR, fiveR, sixR, sevenR, eightR := headR, headR, headR, headR, headR, headR, headR, headR
	oneC, twoC, threeC, fourC, fiveC, sixC, sevenC, eightC := headC, headC, headC, headC, headC, headC, headC, headC

	cmdRgx, _ := regexp.Compile("(\\w)\\s(\\d+)")
	for _, line := range input {

		if cmdRgx.MatchString(line) {

			matches := cmdRgx.FindAllStringSubmatch(line, -1)
			direction := matches[0][1]
			steps, err := strconv.Atoi(matches[0][2])
			if err != nil {
				log.Fatal(err)
			}

			for steps != 0 {

				grid[headR][headC] = "."
				switch direction {
				case right:
					headC += 1
				case left:
					headC -= 1
				case up:
					headR -= 1
				case down:
					headR += 1
				}
				grid[headR][headC] = head

				oneR, oneC = move(one, direction, oneR, oneC, headR, headC, grid)
				twoR, twoC = move(two, direction, twoR, twoC, oneR, oneC, grid)
				threeR, threeC = move(three, direction, threeR, threeC, twoR, twoC, grid)
				fourR, fourC = move(four, direction, fourR, fourC, threeR, threeC, grid)
				fiveR, fiveC = move(five, direction, fiveR, fiveC, fourR, fourC, grid)
				sixR, sixC = move(six, direction, sixR, sixC, fiveR, fiveC, grid)
				sevenR, sevenC = move(seven, direction, sevenR, sevenC, sixR, sixC, grid)
				eightR, eightC = move(eigth, direction, eightR, eightC, sevenR, sevenC, grid)

				tailR, tailC = move(tail, direction, tailR, tailC, eightR, eightC, grid)
				visited[fmt.Sprintf("%d%d", tailR, tailC)] += 1

				steps--
			}
		}
	}
	return fmt.Sprintf("%d", len(visited))
}

func move(name, direction string, currentR, currentC, prevR, prevC int, grid [][]string) (int, int) {

	if !adjacent(currentR, currentC, prevR, prevC) {
		if name == tail {
			grid[currentR][currentC] = "#"
		} else {
			grid[currentR][currentC] = "."
		}
		if diagonally(currentR, currentC, prevR, prevC) {
			//move diagonally towards previous knot

			//if previous above current
			if prevR < currentR {
				currentR -= 1
			} else {
				//else previous below current
				currentR += 1
			}
			//if previous to the left of current
			if prevC < currentC {
				currentC -= 1
			} else {
				currentC += 1
			}
		} else {

			// move along same column
			if currentC == prevC {
				if prevR < currentR {
					currentR -= 1
				} else {
					currentR += 1
				} // move along same row
			} else if currentR == prevR {
				if prevC < currentC {
					currentC -= 1
				} else {
					currentC += 1
				}
			}
		}
		grid[currentR][currentC] = name
	}
	return currentR, currentC
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
