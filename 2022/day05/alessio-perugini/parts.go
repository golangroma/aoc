package main

import (
	"strconv"
	"strings"
)

type Crate struct {
	elems []rune
}

func (s *Crate) Push(v rune) {
	s.elems = append(s.elems, v)
}

func (s *Crate) Pop() rune {
	v := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return v
}

func (s *Crate) PushX(v []rune) {
	s.elems = append(s.elems, v...)
}

func (s *Crate) PopX(x int) []rune {
	v := s.elems[len(s.elems)-x : len(s.elems)]
	s.elems = s.elems[:len(s.elems)-x]
	return v
}

func PartOne(input []string) string {
	createsStartingIndex := getCreatesStartingIndex(input)
	crates := initializeCrates(input[:createsStartingIndex])
	for _, v := range input[createsStartingIndex+1:] {
		nItemsToMove, fromCrate, toCrate := parseMovement(v)
		for i := 0; i < nItemsToMove; i++ {
			crates[toCrate].Push(crates[fromCrate].Pop())
		}
	}
	return getFirstElementOfEachCreates(crates)
}

func PartTwo(input []string) string {
	createsStartingIndex := getCreatesStartingIndex(input)
	crates := initializeCrates(input[:createsStartingIndex])
	for _, v := range input[createsStartingIndex+1:] {
		nItemsToMove, fromCrate, toCrate := parseMovement(v)
		crates[toCrate].PushX(crates[fromCrate].PopX(nItemsToMove))
	}
	return getFirstElementOfEachCreates(crates)
}

func getCreatesStartingIndex(input []string) int {
	for i, v := range input {
		if v == "" {
			return i
		}
	}
	return 0
}

func initializeCrates(input []string) map[int]*Crate {
	crates := make(map[int]*Crate)
	for j := 1; j < len(input[len(input)-1]); j += 4 {
		crates[(j/4)+1] = &Crate{}
	}

	for j := len(input) - 1; j >= 0; j-- {
		for k := 1; k < len(input[j]); k += 4 {
			elm := rune(input[j][k])
			if elm != ' ' {
				crates[(k/4)+1].Push(elm)
			}
		}
	}

	return crates
}

func parseMovement(v string) (int, int, int) {
	movements := strings.Fields(v)
	nItemsToMove, _ := strconv.Atoi(movements[1])
	fromCrate, _ := strconv.Atoi(movements[3])
	toCrate, _ := strconv.Atoi(movements[5])
	return nItemsToMove, fromCrate, toCrate
}

func getFirstElementOfEachCreates(crates map[int]*Crate) string {
	b := &strings.Builder{}
	for i := 1; i <= len(crates); i++ {
		b.WriteRune(crates[i].Pop())
	}
	return b.String()
}
