package main

import (
	"math"
	"strconv"
)

type Pos struct {
	x int
	y int
}

func PartOne(input []string) string {
	result := 0

	grid := make([][]int, 500)

	for i := range grid {
		grid[i] = make([]int, 500)
	}

	startingPos := Pos{250, 250}
	head := Pos{startingPos.x, startingPos.y}
	tail := Pos{startingPos.x, startingPos.y}

	grid[startingPos.x][startingPos.y] = 1

	for _, line := range input {
		direction := line[0]
		distance, _ := strconv.Atoi(line[2:])

		switch direction {
		case 'R':
			for i := 0; i < distance; i++ {
				head.y++
				Follow(head, &tail, &grid, true)
			}
		case 'L':
			for i := 0; i < distance; i++ {
				head.y--
				Follow(head, &tail, &grid, true)
			}
		case 'U':
			for i := 0; i < distance; i++ {
				head.x--
				Follow(head, &tail, &grid, true)
			}
		case 'D':
			for i := 0; i < distance; i++ {
				head.x++
				Follow(head, &tail, &grid, true)
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				result++
			}
		}
	}

	return strconv.Itoa(result)
}

func Follow(head Pos, tail *Pos, grid *[][]int, updateGrid bool) {
	distX := float64(head.x - tail.x)
	distY := float64(head.y - tail.y)
	absDistX := math.Abs(distX)
	absDistY := math.Abs(distY)

	if math.Max(absDistX, absDistY) < 2 {
		return
	}

	if absDistX >= 1 {
		(*tail).x += int(math.Copysign(1, distX))
	}

	if absDistY >= 1 {
		(*tail).y += int(math.Copysign(1, distY))
	}

	if updateGrid {
		(*grid)[tail.x][tail.y]++
	}
}

func PartTwo(input []string) string {
	result := 0

	grid := make([][]int, 500)

	for i := range grid {
		grid[i] = make([]int, 500)
	}

	startingPos := Pos{250, 250}

	nodes := make([]Pos, 10)

	for i := range nodes {
		nodes[i] = startingPos
	}

	grid[startingPos.x][startingPos.y] = 1

	for _, line := range input {
		direction := line[0]
		distance, _ := strconv.Atoi(line[2:])

		switch direction {
		case 'R':
			for i := 0; i < distance; i++ {
				nodes[0].y++

				updateGrid := false

				for j := 0; j < len(nodes)-1; j++ {
					if j == len(nodes)-2 {
						updateGrid = true
					}

					Follow(nodes[j], &nodes[j+1], &grid, updateGrid)
				}
			}
		case 'L':
			for i := 0; i < distance; i++ {
				nodes[0].y--

				updateGrid := false

				for j := 0; j < len(nodes)-1; j++ {
					if j == len(nodes)-2 {
						updateGrid = true
					}

					Follow(nodes[j], &nodes[j+1], &grid, updateGrid)
				}
			}
		case 'U':
			for i := 0; i < distance; i++ {
				nodes[0].x--

				updateGrid := false

				for j := 0; j < len(nodes)-1; j++ {
					if j == len(nodes)-2 {
						updateGrid = true
					}

					Follow(nodes[j], &nodes[j+1], &grid, updateGrid)
				}
			}
		case 'D':
			for i := 0; i < distance; i++ {
				nodes[0].x++

				updateGrid := false
				for j := 0; j < len(nodes)-1; j++ {
					if j == len(nodes)-2 {
						updateGrid = true
					}

					Follow(nodes[j], &nodes[j+1], &grid, updateGrid)
				}
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				result++
			}
		}
	}

	return strconv.Itoa(result)
}
