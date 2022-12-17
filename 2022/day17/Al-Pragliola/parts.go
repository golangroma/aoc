package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type World struct {
	Map      map[int]map[int]byte
	HighestY int
}

func NewWorld() *World {
	return &World{
		Map: map[int]map[int]byte{},
	}
}

func (w *World) String() string {
	var str string

	for y := w.HighestY + 10; y >= 0; y-- {
		for x := 0; x < 7; x++ {

			str += string(w.Get(complex(float64(x), float64(y))))
		}

		str += "\n"
	}

	return str
}

func (w *World) Get(pos complex128) byte {
	x := int(real(pos))
	y := int(imag(pos))

	if x < 0 || x > 6 || y < 0 {
		return '#'
	}

	if w.Map[x] == nil {
		return '.'
	}

	if w.Map[x][y] == 0 {
		return '.'
	}

	return w.Map[x][y]
}

func (w *World) Set(pos complex128, val byte) {
	x := int(real(pos))
	y := int(imag(pos))

	if w.Map[x] == nil {
		w.Map[x] = make(map[int]byte)
	}

	w.Map[x][y] = val
}

type Rock struct {
	Shape []complex128
}

func NewRock(t int, high float64) *Rock {
	switch t {
	case 0: // flat
		return &Rock{
			Shape: []complex128{
				complex(2, high+3),
				complex(3, high+3),
				complex(4, high+3),
				complex(5, high+3),
			},
		}
	case 1: // cross
		return &Rock{
			Shape: []complex128{
				complex(3, high+3),
				complex(2, high+4),
				complex(3, high+4),
				complex(4, high+4),
				complex(3, high+5),
			},
		}
	case 2: // backward L
		return &Rock{
			Shape: []complex128{
				complex(2, high+3),
				complex(3, high+3),
				complex(4, high+3),
				complex(4, high+4),
				complex(4, high+5),
			},
		}
	case 3: // |
		return &Rock{
			Shape: []complex128{
				complex(2, high+3),
				complex(2, high+4),
				complex(2, high+5),
				complex(2, high+6),
			},
		}
	case 4: // box
		return &Rock{
			Shape: []complex128{
				complex(2, high+3),
				complex(3, high+3),
				complex(2, high+4),
				complex(3, high+4),
			},
		}
	default:
		return &Rock{}
	}
}

func (r *Rock) Move(world *World, move complex128) bool {
	var possibleMoves [][]complex128

	for _, s := range r.Shape {
		to := complex(real(s), imag(s)) + move
		toVal := world.Get(to)

		if toVal != '.' && toVal != '@' {
			return false
		}

		possibleMoves = append(possibleMoves, []complex128{s, to})
	}

	for _, m := range possibleMoves {
		world.Set(m[0], '.')
	}

	for i, m := range possibleMoves {
		world.Set(m[1], '@')
		r.Shape[i] = m[1]
	}

	return true
}

func (r *Rock) MoveLeft(world *World) bool {
	return r.Move(world, complex(-1, 0))
}

func (r *Rock) MoveRight(world *World) bool {
	return r.Move(world, complex(1, 0))
}

func (r *Rock) MoveDown(world *World) bool {
	return r.Move(world, complex(0, -1))
}

func PartOne(input []string) string {
	world := NewWorld()
	windJet := input[0]
	maxRocks := 2022

	GetCyclesWithIncrements(world, windJet, maxRocks, true)

	return strconv.Itoa(world.HighestY)
}

func Stringify(in map[int]int) []string {
	var out []string

	for _, v := range in {
		out = append(out, strconv.Itoa(v))
	}

	return out
}

func Sum(incr []int, max int) int {
	var sum int

	for i := 0; i < max; i++ {
		sum += incr[i%len(incr)]
	}

	return sum
}

type Cycles struct {
	StartingRock int
	Length       int
	Height       int
	LatestWind   int
}

func GetCyclesWithIncrements(world *World, windJet string, maxRocks int, skip bool) ([]Cycles, [][]int) {
	var cycles []Cycles
	repeat0 := make(map[int]map[int]int)

	increments := make([][]int, 2)

	k := 0
	j := 0

	for i := 0; i < maxRocks; i++ {
		if i%5 == 0 {
			currentWind := j % len(windJet)
			if repeat0[k] != nil && repeat0[k][currentWind] != 0 {
				if k == 0 {
					cycles = append(cycles, Cycles{
						StartingRock: 0,
						Height:       world.HighestY,
						Length:       len(repeat0[k]),
						LatestWind:   currentWind,
					})
				} else {
					cycles = append(cycles, Cycles{
						StartingRock: cycles[k-1].StartingRock + cycles[k-1].Length,
						Height:       world.HighestY - (k-1)*cycles[k-1].Height - cycles[0].Height,
						Length:       len(repeat0[k]),
						LatestWind:   currentWind,
					})
				}

				k++

				if k > 1 {
					if strings.Join(Stringify(repeat0[k-1]), "") == strings.Join(Stringify(repeat0[k-2]), "") && !skip {
						break
					}
				}
			}

			if repeat0[k] == nil {
				repeat0[k] = make(map[int]int, len(windJet))
			}

			repeat0[k][j%len(windJet)]++
		}

		if k < 2 {
			if increments[k] == nil {
				increments[k] = make([]int, 0)
			}

			increments[k] = append(increments[k], world.HighestY)
		}

		rock := NewRock(i%5, float64(world.HighestY))

		for _, s := range rock.Shape {
			world.Set(s, '@')
		}

		rest := false

		for {
			wind := windJet[j%len(windJet)]

			if wind == '<' {
				rock.MoveLeft(world)
			} else {
				rock.MoveRight(world)
			}

			rest = !rock.MoveDown(world)

			j++

			if rest {
				high := 0

				for _, s := range rock.Shape {
					if int(imag(s)) > high {
						high = int(imag(s))
					}

					world.Set(s, '#')
				}

				high += 1

				if high > world.HighestY {
					world.HighestY = high
				}

				break
			}
		}
	}

	return cycles, increments
}

func GetHighest(cycles []Cycles, increments [][]int, maxRocks int) int {
	div := int(math.Floor(float64((maxRocks - (cycles[0].Length * 5)) / (cycles[len(cycles)-1].Length * 5))))

	highest := (div * cycles[len(cycles)-1].Height) + (cycles[0].Height)

	currentRock := (cycles[0].Length * 5) + (div * cycles[len(cycles)-1].Length * 5)

	if currentRock > maxRocks {
		return highest
	}

	realIncrements := make([]int, len(increments[1]))

	for i := 0; i < len(increments[1]); i++ {
		if i == 0 {
			realIncrements[i] = increments[1][i] - increments[0][len(increments[0])-1]
			continue
		}

		realIncrements[i] = increments[1][i] - increments[1][i-1]
	}

	sum := Sum(realIncrements, maxRocks-currentRock)

	return highest + sum
}

func PartTwo(input []string) string {
	world := NewWorld()
	windJet := input[0]
	maxRocks := 1000000000000

	cycles, increments := GetCyclesWithIncrements(world, windJet, maxRocks, false)

	highest := GetHighest(cycles, increments, maxRocks)

	return fmt.Sprintf("%d", highest)
}
