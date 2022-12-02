package main

import (
	"strconv"
	"strings"
)

type Choice int

const (
	Unknown Choice = iota
	Rock
	Paper
	Scissors
)

func NewChoice(v string) Choice {
	switch v {
	case `A`, `X`:
		return Rock
	case `B`, `Y`:
		return Paper
	case `C`, `Z`:
		return Scissors
	default:
		return Unknown
	}
}

func NewChoiceFromResult(c Choice, v string) Choice {
	switch v {
	case `X`: // lose
		switch c {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}

	case `Y`: // draw
		return c

	case `Z`: // win
		switch c {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}

	return Unknown
}

func (c Choice) Score() int {
	return int(c)
}

type Round struct {
	Player1 Choice
	Player2 Choice
}

func (r *Round) TotalScore() int {
	score := r.Player2.Score()

	// draw
	if r.Player1 == r.Player2 {
		return score + 3
	}

	if r.Player1 == Rock {
		// win
		if r.Player2 == Paper {
			return score + 6
		}
		// lost
		return score
	}

	if r.Player1 == Paper {
		// win
		if r.Player2 == Scissors {
			return score + 6
		}
		// lost
		return score
	}

	if r.Player1 == Scissors {
		// win
		if r.Player2 == Rock {
			return score + 6
		}
		// lost
		return score
	}

	return score
}

func NewRound(ch1, ch2 Choice) *Round {

	return &Round{
		Player1: ch1,
		Player2: ch2,
	}
}

func PartOne(input []string) string {
	rounds := []*Round{}

	for _, line := range input {
		arr := strings.Split(line, " ")
		round := NewRound(NewChoice(arr[0]), NewChoice(arr[1]))
		rounds = append(rounds, round)
	}

	totalScore := 0
	for _, r := range rounds {
		totalScore += r.TotalScore()
	}

	return strconv.Itoa(totalScore)
}

func PartTwo(input []string) string {
	rounds := []*Round{}

	for _, line := range input {
		arr := strings.Split(line, " ")

		ch1 := NewChoice(arr[0])
		round := NewRound(ch1, NewChoiceFromResult(ch1, arr[1]))
		rounds = append(rounds, round)
	}

	totalScore := 0
	for _, r := range rounds {
		totalScore += r.TotalScore()
	}

	return strconv.Itoa(totalScore)
}
