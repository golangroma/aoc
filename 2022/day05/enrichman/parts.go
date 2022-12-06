package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/utils"
)

type Stack struct {
	ID     int
	Crates []*Crate
}

func (s *Stack) AddCrate(v string) {
	crate := Crate(v)
	s.Crates = append(s.Crates, &crate)
}

func (s *Stack) Pick() *Crate {
	top := s.Crates[len(s.Crates)-1]
	s.Crates = s.Crates[:len(s.Crates)-1]
	return top
}

func (s *Stack) Push(c *Crate) {
	s.Crates = append(s.Crates, c)
}

func (s *Stack) String() string {
	var str string
	for _, c := range s.Crates {
		str += fmt.Sprint(string(*c))
	}
	return fmt.Sprintf("Stack{ID: %d, Crates{%s}}", s.ID, str)
}

type Crate string

type Arrangement struct {
	Moves int
	From  int
	To    int
}

type Crane interface {
	Move()
}

type SimpleCrane struct {
	Stacks       []*Stack
	Arrangements []*Arrangement
}

func (sc *SimpleCrane) Move() {
	for _, arrangement := range sc.Arrangements {
		for m := 0; m < arrangement.Moves; m++ {
			crate := sc.Stacks[arrangement.From-1].Pick()
			sc.Stacks[arrangement.To-1].Push(crate)
		}
	}
}

type SuperCrane struct {
	Stacks       []*Stack
	Arrangements []*Arrangement
}

func (sc *SuperCrane) Move() {
	for _, arrangement := range sc.Arrangements {

		ss := sc.Stacks[arrangement.From-1]
		ssTo := sc.Stacks[arrangement.To-1]

		fmt.Println(ss, ssTo)
		fmt.Printf("%+v\n", arrangement)

		crates := ss.Crates[len(ss.Crates)-arrangement.Moves:]
		ss.Crates = ss.Crates[:len(ss.Crates)-arrangement.Moves]
		ssTo.Crates = append(ssTo.Crates, crates...)

		fmt.Println(ss, ssTo)
	}
}

func NewArrangement(s string) *Arrangement {
	conv := func(val string) int {
		i, err := strconv.Atoi(val)
		utils.CheckErr(err)
		return i
	}

	arr := strings.Fields(s)
	if len(arr) > 5 {
		return &Arrangement{
			Moves: conv(arr[1]),
			From:  conv(arr[3]),
			To:    conv(arr[5]),
		}
	}
	return nil
}

func PartOne(input []string) string {
	stackInput, arrangementsInput := splitInput(input)

	stacks := buildStack(stackInput)
	arrangements := buildArrangements(arrangementsInput)

	crane := &SimpleCrane{
		Stacks:       stacks,
		Arrangements: arrangements,
	}
	crane.Move()

	return printTopStacks(crane.Stacks)
}

func PartTwo(input []string) string {
	stackInput, arrangementsInput := splitInput(input)

	stacks := buildStack(stackInput)
	arrangements := buildArrangements(arrangementsInput)

	crane := &SuperCrane{
		Stacks:       stacks,
		Arrangements: arrangements,
	}
	crane.Move()

	return printTopStacks(crane.Stacks)
}

func splitInput(input []string) ([]string, []string) {
	stackInput := []string{}
	arrangementsInput := []string{}

	for i, line := range input {
		if line == "" {
			stackInput = input[0:i]
			arrangementsInput = input[i+1:]
			break
		}
	}

	return stackInput, arrangementsInput
}

func buildStack(stackInput []string) []*Stack {
	printDrawing(stackInput)

	var stacks []*Stack

	// start from the bottom of the drawing stack
	for i := len(stackInput) - 1; i >= 0; i-- {

		// first line is the ID
		if i == len(stackInput)-1 {
			// init stacks
			stacksNum := len(strings.Fields(stackInput[i]))
			stacks = make([]*Stack, stacksNum)
			for s := 0; s < stacksNum; s++ {
				stacks[s] = &Stack{ID: s + 1, Crates: []*Crate{}}
			}
			continue
		}

		stackLine := strings.Fields(stackInput[i])

		for c, crate := range stackLine {
			if crate == "[-]" {
				continue
			}
			stacks[c].AddCrate(crate)
		}
	}

	return stacks
}

func buildArrangements(arrangementsInput []string) []*Arrangement {
	arrangements := []*Arrangement{}

	for _, in := range arrangementsInput {
		arrangement := NewArrangement(in)
		if arrangement == nil {
			continue
		}
		arrangements = append(arrangements, arrangement)
	}

	return arrangements
}

func printDrawing(drawing []string) {
	for i, line := range drawing {
		if strings.HasPrefix(line, "    ") {
			line = strings.Replace(line, "    ", "[-] ", 1)
		}
		line = strings.ReplaceAll(line, "    ", " [-]")

		drawing[i] = line
		fmt.Println(line)
	}
}

func printTopStacks(stacks []*Stack) string {
	builder := strings.Builder{}
	for _, s := range stacks {
		cr := *s.Pick()
		builder.WriteByte([]byte(cr)[1])
	}
	return builder.String()
}
