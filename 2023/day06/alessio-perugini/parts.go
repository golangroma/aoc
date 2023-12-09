package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	time     int
	distance int
}

func Parse(input []string) []Game {
	inputTime := strings.TrimPrefix(input[0], "Time:")
	inputDistance := strings.TrimPrefix(input[1], "Distance:")
	games := []Game{}
	for _, v := range strings.Split(inputTime, " ") {
		if v == "" {
			continue
		}
		games = append(games, Game{time: mustConvertToInt(v)})
	}

	i := 0
	for _, v := range strings.Split(inputDistance, " ") {
		if v == "" {
			continue
		}
		games[i].distance = mustConvertToInt(v)
		i++
	}

	return games
}

func playGame(g Game) int {
	sum := 0
	for i := 0; i < g.time; i++ {
		hold := i
		distance := (g.time - hold) * hold
		if distance > g.distance {
			sum++
		}
	}
	return sum
}

func PartOne(input []string) string {
	games := Parse(input)
	sum := 1
	for _, g := range games {
		if result := playGame(g); result > 0 {
			sum *= result
		}
	}
	return fmt.Sprintf("%v", sum)
}

func PartTwo(input []string) string {
	inputTime := strings.ReplaceAll(strings.TrimPrefix(input[0], "Time:"), " ", "")
	inputDistance := strings.ReplaceAll(strings.TrimPrefix(input[1], "Distance:"), " ", "")
	game := Game{
		time:     mustConvertToInt(inputTime),
		distance: mustConvertToInt(inputDistance),
	}
	sum := 1
	if result := playGame(game); result > 0 {
		sum *= result
	}
	return fmt.Sprintf("%v", sum)
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
