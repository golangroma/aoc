package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Set[T comparable] map[T]bool

// Create a new empty set with the specified initial size.
func NewSet[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

// Add a new key to the set
func (s Set[T]) Add(key T) {
	s[key] = true
}

type Pos struct {
	x, y int
}

type State struct {
	n    int
	rope []Pos
	path Set[Pos]
}

func NewState(n int) *State {
	res := &State{
		n:    n,
		rope: make([]Pos, n),
		path: NewSet[Pos](0),
	}
	res.path.Add(res.rope[n-1])
	return res
}

func (s *State) Move(dir byte) {
	switch dir {
	case 'U':
		s.rope[0].y++
	case 'D':
		s.rope[0].y--
	case 'L':
		s.rope[0].x--
	case 'R':
		s.rope[0].x++
	}
}

func PartOne(input []string) string {
	justString := strings.Join(input, "\n")
	count := run(justString, 2)
	return strconv.Itoa(count)
}

func PartTwo(input []string) string {
	justString := strings.Join(input, "\n")
	count := run(justString, 10)
	return strconv.Itoa(count)
}

func run(input string, n int) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	state := NewState(n)
	for _, line := range lines {
		var dir byte
		var nb int
		fmt.Sscanf(line, "%c %d", &dir, &nb)
		for i := 0; i < nb; i++ {
			state.Move(dir)
			state.MoveTail()
		}
	}
	return len(state.path)
}

func (s *State) MoveTail() {
	for i := 1; i < s.n; i++ {
		delta := Pos{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
		if Abs(delta.x) <= 1 && Abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			s.rope[i].y++
		} else if delta.y < 0 {
			s.rope[i].y--
		}
		if delta.x > 0 {
			s.rope[i].x++
		} else if delta.x < 0 {
			s.rope[i].x--
		}
	}
	s.path.Add(s.rope[s.n-1])
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
