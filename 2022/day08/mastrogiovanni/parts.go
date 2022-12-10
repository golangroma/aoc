package main

import (
	"fmt"
	"strconv"
)

type Grid struct {
	Width  int
	Height int
	Data   [][]int
}

type Direction struct {
	X  int
	Y  int
	Dx int
	Dy int
}

func (grid *Grid) ScenicScore(x, y int) int {
	if x == 0 || y == 0 || x == grid.Width-1 || y == grid.Height-1 {
		return 0
	}

	tot := 1

	// Left
	length := 0
	for x_ := x - 1; x_ >= 0; x_-- {
		length++
		if grid.Data[y][x_] >= grid.Data[y][x] {
			break
		}
	}
	tot *= length

	// Right
	length = 0
	for x_ := x + 1; x_ < grid.Width; x_++ {
		length++
		if grid.Data[y][x_] >= grid.Data[y][x] {
			break
		}
	}
	tot *= length

	// Top
	length = 0
	for y_ := y - 1; y_ >= 0; y_-- {
		length++
		if grid.Data[y_][x] >= grid.Data[y][x] {
			break
		}
	}
	tot *= length

	// Bottom
	length = 0
	for y_ := y + 1; y_ < grid.Height; y_++ {
		length++
		if grid.Data[y_][x] >= grid.Data[y][x] {
			break
		}
	}
	tot *= length
	return tot
}

func (grid *Grid) IsReachableFrom(x, y int, direction Direction) bool {
	x_ := direction.X
	y_ := direction.Y
	for {
		if x_ == x && y_ == y {
			return true
		}
		if grid.Data[y_][x_] >= grid.Data[y][x] {
			return false
		}
		x_ += direction.Dx
		y_ += direction.Dy
	}
}

func (grid *Grid) IsReachable(x_, y_ int) bool {
	if grid.IsReachableFrom(x_, y_, Direction{
		X:  x_,
		Y:  0,
		Dx: 0,
		Dy: 1,
	}) {
		return true
	}
	if grid.IsReachableFrom(x_, y_, Direction{
		X:  x_,
		Y:  grid.Height - 1,
		Dx: 0,
		Dy: -1,
	}) {
		return true
	}
	if grid.IsReachableFrom(x_, y_, Direction{
		X:  0,
		Y:  y_,
		Dx: 1,
		Dy: 0,
	}) {
		return true
	}
	if grid.IsReachableFrom(x_, y_, Direction{
		X:  grid.Width - 1,
		Y:  y_,
		Dx: -1,
		Dy: 0,
	}) {
		return true
	}
	return false
}

func (grid *Grid) Check() int {
	tot := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.IsReachable(x, y) {
				tot++
			}
		}
	}
	return tot
}

func (grid *Grid) BestScenic() int {
	max := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			score := grid.ScenicScore(x, y)
			if score > max {
				max = score
			}
		}
	}
	return max
}

func Parse(inputs []string) Grid {
	width := len(inputs[0])
	height := len(inputs)
	grid := Grid{
		Width:  width,
		Height: height,
		Data:   make([][]int, 0),
	}
	for y := 0; y < height; y++ {
		grid.Data = append(grid.Data, make([]int, 0))
		for x := 0; x < width; x++ {
			n, _ := strconv.Atoi(fmt.Sprintf("%c", inputs[y][x]))
			grid.Data[y] = append(grid.Data[y], n)
		}
	}
	return grid
}

func PartOne(inputs []string) string {
	grid := Parse(inputs)
	return fmt.Sprintf("%d", grid.Check())
}

func PartTwo(inputs []string) string {
	grid := Parse(inputs)
	return fmt.Sprintf("%d", grid.BestScenic())
}
