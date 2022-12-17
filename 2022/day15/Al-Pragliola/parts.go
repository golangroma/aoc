package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

type Range struct {
	start int
	end   int
}

type Pair struct {
	Sensor complex128
	Beacon complex128
}

func GetAllPointsFromAToBAtY(x1, y1, x2, y2, y int, set *map[string]bool) {
	dist := GetManhattanDistance(x1, y1, x2, y2)

	if y < y1-dist || y > y1+dist {
		return
	}

	distFromY := 0

	if y < y1 {
		distFromY = int(math.Abs(float64(y1 - dist - y)))
	} else {
		distFromY = int(math.Abs(float64(y1 + dist - y)))
	}

	for k := x1 - distFromY; k <= x1+distFromY; k++ {
		if (k == x1 && y == y1) || (k == x2 && y == y2) {
			continue
		}

		(*set)[fmt.Sprintf("%d,%d", k, y)] = true
	}
}

func GetRangesAtY(x1, y1, x2, y2, y, max int) (bool, Range) {
	startX, endX := 0, 0
	dist := GetManhattanDistance(x1, y1, x2, y2)

	if y < y1-dist || y > y1+dist {
		return false, Range{startX, endX}
	}

	distFromY := 0

	if y < y1 {
		distFromY = int(math.Abs(float64(y1 - dist - y)))
	} else {
		distFromY = int(math.Abs(float64(y1 + dist - y)))
	}

	return true, Range{int(math.Max(0, float64(x1-distFromY))), int(math.Min(float64(x1+distFromY), float64(max)))}
}

func GetManhattanDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

func FillRanges(pairs []Pair, y, max int) []Range {
	var ranges []Range

	for i := 0; i < len(pairs); i++ {
		ok, r := GetRangesAtY(
			int(real(pairs[i].Sensor)),
			int(imag(pairs[i].Sensor)),
			int(real(pairs[i].Beacon)),
			int(imag(pairs[i].Beacon)),
			y,
			max,
		)

		if ok {
			ranges = append(ranges, r)
		}
	}

	return ranges
}

func FindGapInRanges(ranges []Range, max int) []int {
	if len(ranges) == 1 {
		if ranges[0].start == 0 && ranges[0].end == max {
			return nil
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	if ranges[0].start != 0 {
		return []int{ranges[0].start, ranges[0].start}
	}

	maxR := 0

	for i := 0; i < len(ranges)-1; i++ {
		if ranges[i].end > maxR {
			maxR = ranges[i].end
		}

		if maxR == max {
			return nil
		}

		if maxR+1 < ranges[i+1].start {
			return []int{maxR + 1, maxR + 1}
		}
	}

	return nil
}

func GetPairs(input []string) []Pair {
	reg := regexp.MustCompile(`(-?\d+)`)
	pairs := make([]Pair, 0)
	for i := 0; i < len(input); i++ {
		points := reg.FindAllString(input[i], -1)

		x1, _ := strconv.ParseInt(points[0], 10, 0)
		y1, _ := strconv.ParseInt(points[1], 10, 0)
		x2, _ := strconv.ParseInt(points[2], 10, 0)
		y2, _ := strconv.ParseInt(points[3], 10, 0)

		s := complex(float64(x1), float64(y1))
		b := complex(float64(x2), float64(y2))

		pairs = append(pairs, Pair{s, b})
	}

	return pairs
}

func PartOne(input []string) string {
	y := 2000000

	pairs := GetPairs(input)
	set := make(map[string]bool)

	for i := 0; i < len(pairs); i++ {
		GetAllPointsFromAToBAtY(
			int(real(pairs[i].Sensor)),
			int(imag(pairs[i].Sensor)),
			int(real(pairs[i].Beacon)),
			int(imag(pairs[i].Beacon)),
			y,
			&set,
		)
	}

	return strconv.Itoa(len(set))
}

func PartTwo(input []string) string {
	max := 4000000

	pairs := GetPairs(input)

	for i := 0; i <= max; i++ {
		ranges := FillRanges(pairs, i, max)

		gap := FindGapInRanges(ranges, max)

		if gap != nil {
			return strconv.Itoa(gap[0]*4000000 + i)
		}
	}

	return "not found"
}
