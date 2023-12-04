package main

import (
	"log"
	"strconv"
	"strings"
)

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Reds   int
	Greens int
	Blues  int
}

func (g Game) IsValid(reds, greens, blues int) bool {
	for _, set := range g.Sets {
		if set.Reds > reds {
			return false
		}
		if set.Greens > greens {
			return false
		}
		if set.Blues > blues {
			return false
		}
	}
	return true
}

func (g Game) MinimumCubes() (int, int, int) {
	var reds, greens, blues int

	for _, set := range g.Sets {
		if reds < set.Reds {
			reds = set.Reds
		}
		if greens < set.Greens {
			greens = set.Greens
		}
		if blues < set.Blues {
			blues = set.Blues
		}
	}

	return reds, greens, blues
}

func PartOne(input []string) string {
	sum := 0

	for _, line := range input {
		game := mustParseGame(line)
		if game.IsValid(12, 13, 14) {
			sum += game.ID
		}
	}

	return strconv.Itoa(sum)
}

func PartTwo(input []string) string {
	sum := 0

	for _, line := range input {
		game := mustParseGame(line)
		r, g, b := game.MinimumCubes()
		sum += (r * g * b)
	}

	return strconv.Itoa(sum)
}

func mustParseGame(input string) Game {
	splitted := strings.Split(input, ":")

	gameString := splitted[0]
	idString := strings.Split(gameString, " ")[1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal(err)
	}

	setsString := splitted[1]
	sets := mustParseSets(setsString)

	return Game{
		ID:   id,
		Sets: sets,
	}
}

func mustParseSets(input string) []Set {
	sets := []Set{}

	input = strings.TrimSpace(input)
	setsString := strings.Split(input, ";")

	for _, setString := range setsString {
		set := Set{}

		// setString is like '1 blue, 2 green'
		setString = strings.TrimSpace(setString)
		cubesString := strings.Split(setString, ",")

		for _, cubeString := range cubesString {
			// cubeString is like '1 blue'
			cubeString = strings.TrimSpace(cubeString)
			cubes := strings.Split(cubeString, " ")

			numOfCubes, err := strconv.Atoi(cubes[0])
			if err != nil {
				log.Fatal(err)
			}
			switch cubes[1] {
			case "red":
				set.Reds = numOfCubes
			case "green":
				set.Greens = numOfCubes
			case "blue":
				set.Blues = numOfCubes
			}
		}

		sets = append(sets, set)
	}

	return sets
}
