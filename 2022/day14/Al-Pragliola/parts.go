package main

import (
	"strconv"
	"strings"
)

type World struct {
	Map        map[int]map[int]byte
	MaxX       int
	MaxY       int
	MinX       int
	SandOrigin int
}

func ScanLine(line string, w *World) (int, int, int) {
	maxX := 0
	maxY := 0
	minX := 1<<63 - 10

	steps := strings.Split(line, " -> ")

	for i := 0; i < len(steps)-1; i++ {
		coordsA := strings.Split(steps[i], ",")
		coordsB := strings.Split(steps[i+1], ",")

		xA, _ := strconv.Atoi(coordsA[0])
		yA, _ := strconv.Atoi(coordsA[1])

		if xA > maxX {
			maxX = xA
		}

		if yA > maxY {
			maxY = yA
		}

		if xA < minX {
			minX = xA
		}

		xB, _ := strconv.Atoi(coordsB[0])
		yB, _ := strconv.Atoi(coordsB[1])

		if xB > maxX {
			maxX = xB
		}

		if yB > maxY {
			maxY = yB
		}

		if xB < minX {
			minX = xB
		}

		if xA == xB {
			if yA < yB {
				for y := yA; y <= yB; y++ {
					w.Map[xA][y] = '#'
				}
			} else {
				for y := yB; y <= yA; y++ {
					w.Map[xA][y] = '#'
				}
			}
		} else {
			if xA < xB {
				for x := xA; x <= xB; x++ {
					w.Map[x][yA] = '#'
				}
			} else {
				for x := xB; x <= xA; x++ {
					w.Map[x][yA] = '#'
				}
			}
		}
	}

	return maxX, maxY, minX
}

func NewWorld(input []string, sandOrigin int) *World {
	var w World

	w.MaxX = 0
	w.MaxY = 0
	w.MinX = 1<<63 - 10
	w.SandOrigin = sandOrigin

	w.Map = make(map[int]map[int]byte, 1000)

	for i := 0; i < 1000; i++ {
		w.Map[i] = make(map[int]byte, 1000)
		for j := 0; j < 1000; j++ {
			w.Map[i][j] = '.'
		}
	}

	w.Map[sandOrigin][0] = '+'

	for _, line := range input {
		lmaxX, lmaxY, lminX := ScanLine(line, &w)

		if lmaxX > w.MaxX {
			w.MaxX = lmaxX
		}

		if lmaxY > w.MaxY {
			w.MaxY = lmaxY
		}

		if lminX < w.MinX {
			w.MinX = lminX
		}
	}

	return &w
}

func (w *World) String() string {
	var s string

	for y := 0; y <= w.MaxY; y++ {
		for x := w.MinX; x <= w.MaxX; x++ {
			s += string(w.Map[x][y])
		}
		s += "\n"
	}

	return s
}

func (w *World) AddFloor() {
	for x := 0; x < len(w.Map[w.MaxY+2]); x++ {
		w.Map[x][w.MaxY+2] = '#'
	}

	w.MaxY += 2
}

func (w *World) SimulateDrop(stopFunc func(x, y int, wld *World) bool) (bool, int) {
	var cycles int

	stop := false
	dropX := w.SandOrigin
	dropY := 0

	for {
		if stopFunc(dropX, dropY, w) {
			stop = true
			break
		}

		if w.Map[dropX][dropY+1] == '.' {
			dropY++
			cycles++
			continue
		}

		if w.Map[dropX][dropY+1] == '#' || w.Map[dropX][dropY+1] == 'o' {
			if w.Map[dropX-1][dropY+1] == '.' {
				dropX--
				dropY++
				cycles++
				continue
			}

			if w.Map[dropX+1][dropY+1] == '.' {
				dropX++
				dropY++
				cycles++
				continue
			}

			w.Map[dropX][dropY] = 'o'
			break
		}
	}

	return stop, cycles
}

func PartOne(input []string) string {
	var drops int
	w := NewWorld(input, 500)

	fellCheck := func(x, y int, wld *World) bool {
		if x < w.MinX || x > w.MaxX || y > w.MaxY {
			return true
		}

		return false
	}

	for {
		fellOutside, _ := w.SimulateDrop(fellCheck)
		if fellOutside {
			break
		}
		drops++
	}

	return strconv.Itoa(drops)
}

func PartTwo(input []string) string {
	var drops int
	w := NewWorld(input, 500)

	w.AddFloor()

	reachTopCheck := func(x, y int, wld *World) bool {
		if wld.Map[w.SandOrigin][0] == 'o' {
			return true
		}

		return false
	}

	for {
		topReached, _ := w.SimulateDrop(reachTopCheck)
		if topReached {
			break
		}
		drops++
	}

	return strconv.Itoa(drops)
}
