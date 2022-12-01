package main

import (
	"strconv"
	"strings"

	"github.com/golangroma/aoc/utils"
)

type Command struct {
	Direction Direction
	Units     int
}

type Direction string

const (
	Up      Direction = "up"
	Down    Direction = "down"
	Forward Direction = "forward"
)

type Submarine1 struct {
	Horizontal int
	Depth      int
}

func (s *Submarine1) Move(cmd Command) {
	switch cmd.Direction {
	case Up:
		s.Depth -= cmd.Units
	case Down:
		s.Depth += cmd.Units
	case Forward:
		s.Horizontal += cmd.Units
	}
}

func PartOne(input []string) string {
	commands, err := utils.Convert(input, CommandDonverter)
	utils.CheckErr(err)

	submarine := &Submarine1{}
	for _, cmd := range commands {
		submarine.Move(cmd)
	}

	return strconv.Itoa(submarine.Horizontal * submarine.Depth)
}

func CommandDonverter(lines []string) ([]Command, error) {
	commands := []Command{}

	for _, line := range lines {
		arr := strings.Split(line, " ")
		units, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, err
		}

		commands = append(commands, Command{
			Direction: Direction(arr[0]),
			Units:     units,
		})
	}

	return commands, nil
}

type Submarine2 struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (s *Submarine2) Move(cmd Command) {
	switch cmd.Direction {
	case Up:
		s.Aim -= cmd.Units
	case Down:
		s.Aim += cmd.Units
	case Forward:
		s.Horizontal += cmd.Units
		s.Depth += s.Aim * cmd.Units
	}
}

func PartTwo(input []string) string {
	commands, err := utils.Convert(input, CommandDonverter)
	utils.CheckErr(err)

	submarine := &Submarine2{}
	for _, cmd := range commands {
		submarine.Move(cmd)
	}

	return strconv.Itoa(submarine.Horizontal * submarine.Depth)
}
