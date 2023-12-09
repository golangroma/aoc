package main

import (
	"fmt"
	"strings"
)

const (
	Left  = "L"
	Right = "R"
)

type Game struct {
	nodes map[string]Node
	steps string
}

type Node struct {
	name  string
	left  string
	right string
}

func Parse(input []string) Game {
	g := Game{steps: input[0], nodes: make(map[string]Node)}
	for _, v := range input[2:] {
		curr, next, _ := strings.Cut(v, " = ")
		left, right, _ := strings.Cut(next[1:len(next)-1], ", ")
		g.nodes[curr] = Node{
			name:  curr,
			left:  left,
			right: right,
		}
	}
	return g
}

func PartOne(input []string) string {
	g := Parse(input)
	return fmt.Sprintf("%d", game2("AAA", g, func(s string) bool { return s == "ZZZ" }))
}

func game2(start string, g Game, cmp func(string) bool) int {
	i, steps := 0, 0
	next := start
	for {
		steps++
		if string(g.steps[i]) == Left {
			next = g.nodes[next].left
		} else {
			next = g.nodes[next].right
		}
		if cmp(next) {
			break
		}
		if i+1 < len(g.steps) {
			i++
		} else {
			i = 0
		}
	}
	return steps
}

func PartTwo(input []string) string {
	g := Parse(input)
	next := []string{}
	for k := range g.nodes {
		if strings.HasSuffix(k, "A") {
			next = append(next, k)
		}
	}
	steps := []int{}
	for _, v := range next {
		c := game2(v, g, func(s string) bool { return strings.HasSuffix(s, "Z") })
		steps = append(steps, c)
	}
	return fmt.Sprintf("%d", LCM(steps[0], steps[1], steps[2:]...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
