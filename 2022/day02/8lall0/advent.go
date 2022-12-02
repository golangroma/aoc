package main

import (
	"bufio"
	"fmt"
	"io"
)

const (
	WIN int = iota
	TIE
	LOSE
)

const (
	ROCK int = iota
	PAPER
	SCISSOR
)

var opponent = map[string]int{
	"A": ROCK,    // Sasso
	"B": PAPER,   // Carta
	"C": SCISSOR, // Forbici
}

var score = map[int]int{
	WIN:  6,
	TIE:  3,
	LOSE: 0,
}

var typeScore = map[int]int{
	ROCK:    1,
	PAPER:   2,
	SCISSOR: 3,
}

var results = [][]int{
	// ROCK PAPER SCISSOR
	{TIE, LOSE, WIN}, // ROCK
	{WIN, TIE, LOSE}, // PAPER
	{LOSE, WIN, TIE}, // SCISSOR
}

func part1(r io.Reader) (int, error) {
	mapResults := map[string]int{
		"X": ROCK,    // Sasso
		"Y": PAPER,   // Carta
		"Z": SCISSOR, // Forbici
	}

	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		txt := scanner.Text()
		var a, b string
		_, err := fmt.Sscanf(txt, "%s %s", &a, &b)
		if err != nil {
			return 0, err
		}

		moveA := opponent[a]
		moveB := mapResults[b]

		result := results[moveB][moveA]

		sum += score[result] + typeScore[moveB]
	}

	return sum, nil
}

func part2(r io.Reader) (int, error) {
	mapResults := map[string]int{
		"X": LOSE, // Sasso
		"Y": TIE,  // Carta
		"Z": WIN,  // Forbici
	}

	scanner := bufio.NewScanner(r)

	sum := 0

	for scanner.Scan() {
		txt := scanner.Text()
		var a, b string
		_, err := fmt.Sscanf(txt, "%s %s", &a, &b)
		if err != nil {
			return 0, err
		}

		moveA := opponent[a]
		result := mapResults[b]

		ndx := 0
		for i, row := range results {
			if row[moveA] == result {
				ndx = i
				break
			}
		}

		sum += score[result] + typeScore[ndx]
	}

	return sum, nil
}
