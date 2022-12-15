package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Record struct {
	s Point
	b Point
	d int
}

type Problem struct {
	r    []Record
	minX int
	maxX int
	minY int
	maxY int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(a, b Point) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

func parse(inputs []string) *Problem {
	problem := &Problem{}
	res := make([]Record, len(inputs))
	first := true
	for i, row := range inputs {
		// fmt.Println(row)
		comps := strings.Split(row, " ")
		x, _ := strconv.Atoi(comps[2][2 : len(comps[2])-1])
		y, _ := strconv.Atoi(comps[3][2 : len(comps[3])-1])
		if first {
			problem.minX = x
			problem.maxX = x
			problem.minY = y
			problem.maxY = y
			first = false
		}

		if x < problem.minX {
			problem.minX = x
		}
		if x > problem.maxX {
			problem.maxX = x
		}
		if y < problem.minY {
			problem.minY = y
		}
		if y > problem.maxY {
			problem.maxY = y
		}

		s := Point{x, y}
		bx, _ := strconv.Atoi(comps[8][2 : len(comps[8])-1])
		by, _ := strconv.Atoi(comps[9][2:len(comps[9])])

		if bx < problem.minX {
			problem.minX = bx
		}
		if bx > problem.maxX {
			problem.maxX = bx
		}
		if by < problem.minY {
			problem.minY = by
		}
		if by > problem.maxY {
			problem.maxY = by
		}

		b := Point{bx, by}
		res[i] = Record{s, b, ManhattanDistance(s, b)}
	}
	problem.r = res
	return problem
}

func IsNotInRange(p Point, problem *Problem) bool {
	for _, r := range problem.r {
		if r.b.x == p.x && r.b.y == p.y {
			continue
		}
		dd := ManhattanDistance(r.s, p)
		if dd <= r.d {
			return true
		}
	}
	return false
}

// To improve skip
func CheckBeaconsAndSkip(p Point, problem *Problem) (bool, int) {
	skip := -1
	inside := false
	for _, r := range problem.r {
		if r.b.x == p.x && r.b.y == p.y {
			continue
		}
		dd := ManhattanDistance(r.s, p)
		if dd <= r.d {
			// fmt.Printf("%+v is at distance %d from %+v (closer than %d)\n", p, dd, r.b, r.d)
			inside = true
			if r.s.x <= p.x {
				if skip < r.d-dd {
					skip = r.d - dd
				}
			}
			if r.s.x > p.x {
				max := Abs(r.s.x-p.x) + r.d - Abs(r.s.y-p.y)
				if skip < max {
					skip = max
				}
			}
		}
	}
	return inside, skip
	// fmt.Printf("CheckBeaconsAndSkip: %+v, %+v, %d\n", problem, inside, skip)
	// return inside, skip
}

func PartOne(inputs []string, y int) string {
	p := parse(inputs)
	count := 0
	for x := p.minX; x <= p.maxX; x++ {
		g := Point{x, y}
		found := IsNotInRange(g, p)
		if found {
			count++
		}
	}
	covered := false
	for x := p.minX - 1; !covered; x-- {
		g := Point{x, y}
		found := IsNotInRange(g, p)
		if found {
			count++
		} else {
			covered = true
		}
	}
	covered = false
	for x := p.maxX + 1; !covered; x++ {
		g := Point{x, y}
		found := IsNotInRange(g, p)
		if found {
			count++
		} else {
			covered = true
		}
	}
	return fmt.Sprintf("%d", count)
}

func PartTwo(inputs []string, width int) string {
	p := parse(inputs)
	best := Point{}
	for y := 0; y <= width; y++ {
		for x := 0; x <= width; x++ {
			g := Point{x, y}
			inside, skip := CheckBeaconsAndSkip(g, p)
			if inside {
				if skip > 0 {
					x += skip
					if x > width+1 {
						x = width + 1
					}
				}
			} else {
				best = g
			}
		}
	}
	return fmt.Sprintf("%d", best.x*4000000+best.y)
}
