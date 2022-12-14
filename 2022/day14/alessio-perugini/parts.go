package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Obstacle int

const (
	Air Obstacle = iota
	Rock
	Sand
)

type P struct {
	X, Y int
}

func (p P) Same(p2 P) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p P) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Scan struct {
	cave                 map[P]Obstacle
	yLimit               int
	sandPoured           int
	sandStartingPosition P
}

func NewScan(sandPosition P) *Scan {
	return &Scan{
		cave:                 make(map[P]Obstacle),
		sandStartingPosition: sandPosition,
	}
}

func (m *Scan) Simulate() int {
	for {
		if !m.pourSand(m.sandStartingPosition) {
			return m.sandPoured
		}
	}
}

func (m *Scan) pourSand(currPosition P) bool {
	for {
		if currPosition.Y > m.yLimit {
			return false
		}
		if m.cave[currPosition] == Air {
			currPosition.Y++
			continue
		}

		downLeft := currPosition
		downLeft.X--
		if m.cave[downLeft] == Air {
			currPosition = downLeft
			continue
		}

		downRight := currPosition
		downRight.X++
		if m.cave[downRight] == Air {
			currPosition = downRight
			continue
		}

		currPosition.Y--
		m.cave[currPosition] = Sand
		m.sandPoured++

		if currPosition.Same(m.sandStartingPosition) {
			return false
		}

		break
	}

	return true
}

func (m *Scan) InsertRocks(rocks []P) {
	if len(rocks) == 0 {
		return
	}
	if len(rocks) == 1 {
		m.cave[rocks[0]] = 1
	}

	for i := 1; i < len(rocks); i++ {
		r1, r2 := rocks[i-1], rocks[i]

		m.insertX(r1, r1.X-r2.X)
		m.insertY(r1, r1.Y-r2.Y)

		m.updateMapLimit(r1.Y)
		m.updateMapLimit(r2.Y)
	}
}

func (m *Scan) updateMapLimit(y int) {
	if y > m.yLimit {
		m.yLimit = y
	}
}

func (m *Scan) insertX(r P, diffX int) {
	offset := 1
	if diffX >= 0 {
		offset = -1
	}

	m.cave[r] = Rock
	for i := 0; i < sign(diffX); i++ {
		r.X += offset
		m.cave[r] = Rock
	}
}

func (m *Scan) insertY(r P, diffY int) {
	offset := 1
	if diffY >= 0 {
		offset = -1
	}

	m.cave[r] = Rock
	for i := 0; i < sign(diffY); i++ {
		r.Y += offset
		m.cave[r] = Rock
	}
}

func sign(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func parseInput(v string) []P {
	points := strings.Split(v, " -> ")
	rocks := make([]P, len(points))
	for i := 0; i < len(rocks); i++ {
		rocks[i] = parseRock(points[i])
	}
	return rocks
}

func parseRock(v string) P {
	coords := strings.Split(v, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return P{X: x, Y: y}
}

func PartOne(input []string) string {
	m := NewScan(P{X: 500, Y: 0})
	for _, v := range input {
		m.InsertRocks(parseInput(v))
	}

	return fmt.Sprintf("%d", m.Simulate())
}

func PartTwo(input []string) string {
	sand := P{X: 500, Y: 0}
	m := NewScan(sand)

	for _, v := range input {
		m.InsertRocks(parseInput(v))
	}

	m.yLimit = m.yLimit + 2
	floorRockLeft := P{X: sand.X - m.yLimit - 1, Y: m.yLimit}
	floorRockRight := P{X: sand.X + m.yLimit + 1, Y: m.yLimit}

	m.InsertRocks([]P{floorRockLeft, floorRockRight})

	return fmt.Sprintf("%d", m.Simulate())
}
