package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func move(point *Point, direction string) {
	if direction == "R" {
		point.X = point.X + 1
	}
	if direction == "L" {
		point.X = point.X - 1
	}
	if direction == "U" {
		point.Y = point.Y - 1
	}
	if direction == "D" {
		point.Y = point.Y + 1
	}
}

func moveExtended(tail *Point, head *Point, headPrec *Point) {
	dX := head.X - tail.X
	dY := head.Y - tail.Y
	_dx := 0
	if dX > 0 {
		_dx = 1
	} else if dX < 0 {
		_dx = -1
	}
	_dy := 0
	if dY > 0 {
		_dy = 1
	} else if dY < 0 {
		_dy = -1
	}
	tail.X += _dx
	tail.Y += _dy
}

func moveTail(tail *Point, head *Point, direction string) {
	if direction == "R" {
		tail.X = head.X - 1
		tail.Y = head.Y
	}
	if direction == "L" {
		tail.X = head.X + 1
		tail.Y = head.Y
	}
	if direction == "U" {
		tail.X = head.X
		tail.Y = head.Y + 1
	}
	if direction == "D" {
		tail.X = head.X
		tail.Y = head.Y - 1
	}
}

func isClose(a Point, b Point) bool {
	if math.Abs(float64(a.X-b.X)) > 1 {
		return false
	}
	if math.Abs(float64(a.Y-b.Y)) > 1 {
		return false
	}
	return true
}

func Dump(knots []Point) {
	minX := -10 // knots[0].X
	maxX := 10  // knots[0].X
	minY := -10 // knots[0].Y
	maxY := 10  // knots[0].Y
	for _, p := range knots {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	// minX -= 3
	// maxX += 3
	// minY -= 3
	// maxY += 3
	fmt.Printf("%s\n", knots[0])
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			found := false
			for i, p := range knots {
				if p.X == x && p.Y == y {
					fmt.Printf("%d", i)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}

func PartOne(input []string) string {
	visited := make(map[string]struct{})
	head := Point{0, 4}
	tail := Point{0, 4}
	for _, command := range input {
		comps := strings.Split(command, " ")
		length, _ := strconv.Atoi(comps[1])
		for i := 0; i < length; i++ {
			move(&head, comps[0])
			if !isClose(head, tail) {
				moveTail(&tail, &head, comps[0])
			}
			visited[tail.String()] = struct{}{}
		}
	}
	return fmt.Sprintf("%d", len(visited))
}

func PartTwo(input []string) string {
	visited := make(map[string]struct{})
	knots := []Point{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	for _, command := range input {
		comps := strings.Split(command, " ")
		length, _ := strconv.Atoi(comps[1])
		for i := 0; i < length; i++ {
			headPrec := knots[0]
			move(&knots[0], comps[0])
			for j := 1; j < len(knots); j++ {
				if !isClose(knots[j-1], knots[j]) {
					moveExtended(&knots[j], &knots[j-1], &headPrec)
					headPrec = knots[j]
				} else {
					break
				}
			}
			visited[knots[len(knots)-1].String()] = struct{}{}
		}
	}
	return fmt.Sprintf("%d", len(visited))
}
