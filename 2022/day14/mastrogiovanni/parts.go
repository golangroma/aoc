package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Cell int

const (
	Empty = iota
	Wall
	Sand
)

type Grid struct {
	minX   int
	maxX   int
	maxY   int
	width  int
	height int
	data   [][]Cell
}

func parse_size(inputs []string, grid *Grid) {
	minX := 500
	maxX := 500
	maxY := 0
	for _, row := range inputs {
		comps := strings.Split(row, " -> ")
		for _, comp := range comps {
			coords := strings.Split(comp, ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	grid.maxX = maxX
	grid.maxY = maxY
	grid.minX = minX
}

func parse_size2(inputs []string, grid *Grid) {
	maxY := 0
	for _, row := range inputs {
		comps := strings.Split(row, " -> ")
		for _, comp := range comps {
			coords := strings.Split(comp, ",")
			// x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			if y > maxY {
				maxY = y
			}
		}
	}
	grid.maxY = maxY + 2
	grid.maxX = 500 + grid.maxY*3
	grid.minX = 500 - grid.maxY*3
}

func parse_empty_grid(grid *Grid) {
	grid.width = grid.maxX - grid.minX + 1
	grid.height = grid.maxY + 1
	grid.data = make([][]Cell, grid.height)
	for i := 0; i < grid.height; i++ {
		grid.data[i] = make([]Cell, grid.width)
	}
}

func (g *Grid) dump() {
	for i := 0; i < len(g.data); i++ {
		fmt.Println()
		for j := 0; j < len(g.data[0]); j++ {
			if g.data[i][j] == Empty {
				fmt.Print(".")
			} else if g.data[i][j] == Wall {
				fmt.Print("#")
			} else if g.data[i][j] == Sand {
				fmt.Print("o")
			}
		}
	}
	fmt.Println()
}

func fill_last_line(grid *Grid) {
	for j := 0; j < grid.width; j++ {
		grid.data[grid.maxY][j] = Wall
	}
}

func fill_the_grid(inputs []string, grid *Grid) {
	for _, row := range inputs {
		segment := make([]string, 0)
		comps := strings.Split(row, " -> ")
		for _, comp := range comps {
			if len(segment) == 1 {
				startcoords := strings.Split(segment[0], ",")
				startX, _ := strconv.Atoi(startcoords[0])
				startY, _ := strconv.Atoi(startcoords[1])

				coords := strings.Split(comp, ",")
				endX, _ := strconv.Atoi(coords[0])
				endY, _ := strconv.Atoi(coords[1])

				// Draw line
				dx := 0
				dy := 0
				if startX > endX {
					dx = -1
				}
				if startX < endX {
					dx = 1
				}
				if startY > endY {
					dy = -1
				}
				if startY < endY {
					dy = 1
				}
				// fmt.Printf("(%d,%d) => (%d,%d): %d %d\n", startX, startY, endX, endY, dx, dy)
				x := startX
				y := startY
				for {
					grid.data[y][x-grid.minX] = Wall
					if x == endX && y == endY {
						break
					}
					x += dx
					y += dy
				}
				segment[0] = comp
			} else {
				segment = append(segment, comp)
			}

		}
	}
}

func parse(inputs []string) *Grid {
	grid := &Grid{}
	parse_size(inputs, grid)
	parse_empty_grid(grid)
	fill_the_grid(inputs, grid)
	return grid
}

func parse2(inputs []string) *Grid {
	grid := &Grid{}
	parse_size2(inputs, grid)
	parse_empty_grid(grid)
	fill_the_grid(inputs, grid)
	fill_last_line(grid)
	return grid
}

func simulator(grid *Grid) bool {
	x := 500 - grid.minX
	y := 0
	for {
		if grid.data[0][500-grid.minX] == Sand {
			return false
		}
		if y+1 > grid.height-1 {
			return false
		}
		if grid.data[y+1][x] == Empty {
			y++
			continue
		}
		if x-1 < 0 {
			return false
		}
		if grid.data[y+1][x-1] == Empty {
			x--
			y++
			continue
		}
		if x+1 > grid.width-1 {
			return false
		}
		if grid.data[y+1][x+1] == Empty {
			x++
			y++
			continue
		}
		grid.data[y][x] = Sand
		break
	}
	return true
}

func PartOne(inputs []string) string {
	grid := parse(inputs)
	// grid.dump()
	for i := 0; ; i++ {
		check := simulator(grid)
		if !check {
			return fmt.Sprintf("%d", i)
		}
		// grid.dump()
	}
}

func PartTwo(inputs []string) string {
	grid := parse2(inputs)
	// grid.dump()
	for i := 0; ; i++ {
		check := simulator(grid)
		if !check {
			return fmt.Sprintf("%d", i)
		}
		// grid.dump()
	}
}
