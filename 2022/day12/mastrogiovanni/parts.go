package main

import "fmt"

type Point struct {
	x int
	y int
}

type HillMap struct {
	start  Point
	stop   Point
	width  int
	height int
	grid   [][]int
}

func Parse(inputs []string) *HillMap {
	width := len(inputs[0])
	height := len(inputs)
	result := &HillMap{}
	result.grid = make([][]int, height)
	result.width = width
	result.height = height
	for i, row := range inputs {
		result.grid[i] = make([]int, width)
		for j := 0; j < width; j++ {
			if row[j] == 'S' {
				result.start = Point{x: j, y: i}
			} else if row[j] == 'E' {
				result.stop = Point{x: j, y: i}
				result.grid[i][j] = int('z' - 'a')
			} else {
				result.grid[i][j] = int(row[j] - 'a')
			}
		}
	}
	return result
}

func Visit(hill_map *HillMap) [][]int {

	stack := make([]Point, 0)

	distances := make([][]int, hill_map.height)
	for i := 0; i < hill_map.height; i++ {
		distances[i] = make([]int, hill_map.width)
		for j := 0; j < hill_map.width; j++ {
			distances[i][j] = -1
		}
	}

	distances[hill_map.stop.y][hill_map.stop.x] = 0
	stack = append(stack, hill_map.stop)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	count := 1

	for len(stack) > 0 {
		item := stack[0]
		stack = stack[1:]

		current := distances[item.y][item.x]
		for i := 0; i < 4; i++ {

			x := item.x + dx[i]
			y := item.y + dy[i]

			if !(x >= 0 && x < hill_map.width && y >= 0 && y < hill_map.height) {
				continue
			}

			if hill_map.grid[item.y][item.x]-hill_map.grid[y][x] <= 1 {
				old := distances[y][x]
				if old == -1 {
					count++
					distances[y][x] = current + 1
					stack = append(stack, Point{x: x, y: y})
				}
			}
		}
	}

	return distances
}

func DumpGrid(d [][]int) {
	for j := 0; j < len(d); j++ {
		for i := 0; i < len(d[j]); i++ {
			fmt.Printf("%3d", d[j][i])
		}
		fmt.Println()
	}
	fmt.Println()
}

func PartOne(inputs []string) string {
	hill_map := Parse(inputs)
	d := Visit(hill_map)
	min_len := d[hill_map.start.y][hill_map.start.x]
	return fmt.Sprintf("%d", min_len)
}

func PartTwo(inputs []string) string {
	hill_map := Parse(inputs)
	d := Visit(hill_map)
	min_len := -1
	for i := 0; i < hill_map.height; i++ {
		for j := 0; j < hill_map.width; j++ {
			if hill_map.grid[i][j] == 0 && d[i][j] >= 0 {
				if min_len == -1 {
					min_len = d[i][j]
				} else if min_len > d[i][j] {
					min_len = d[i][j]
				}
			}
		}
	}
	return fmt.Sprintf("%d", min_len)
}
