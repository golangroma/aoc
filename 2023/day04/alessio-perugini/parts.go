package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	var sum int
	for _, v := range input {
		_, nextToken, _ := strings.Cut(v, ": ")
		winningBoard, myNumbers, _ := strings.Cut(nextToken, " | ")
		wBoard, gNumbers := parseBoard(winningBoard), parseBoard(myNumbers)
		winningNumbers, points := []int{}, 0
		for _, n := range gNumbers {
			if slices.Contains(wBoard, n) {
				winningNumbers = append(winningNumbers, n)
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		//fmt.Println(points)
		sum += points
	}

	return fmt.Sprintf("%v", sum)
}

type Card struct {
	id             int
	boardNumbers   []int
	guessedNumbers []int
	latestID       int
}

func (c Card) Play() []int {
	result := []int{}
	sum := 0
	for _, n := range c.guessedNumbers {
		if slices.Contains(c.boardNumbers, n) {
			sum++
			result = append(result, c.id+sum)
		}
	}
	return result
}

func PartTwo(input []string) string {
	var sum int
	parsedCards := map[int]Card{}
	cardSeen := map[int]int{}
	//latestCardID := len(input)

	queue := []int{}
	for _, v := range input {
		var gameID int
		game, nextToken, _ := strings.Cut(v, ": ")
		_, _ = fmt.Sscanf(game, "Card %d", &gameID)
		winningBoard, myNumbers, _ := strings.Cut(nextToken, " | ")
		wBoard, gNumbers := parseBoard(winningBoard), parseBoard(myNumbers)
		c := Card{
			id:             gameID,
			boardNumbers:   wBoard,
			guessedNumbers: gNumbers,
		}
		parsedCards[gameID] = c
		copies := c.Play()
		if len(copies) > 0 {
			queue = append(queue, c.Play()...)
			sum++
			cardSeen[gameID]++
		}
	}

	for len(queue) > 0 {
		var id int
		id, queue = queue[0], queue[1:]
		copies := parsedCards[id].Play()
		cardSeen[id]++
		sum++
		fmt.Println(id, copies)
		queue = append(queue, copies...)
	}
	fmt.Println(cardSeen)

	return fmt.Sprintf("%v", sum)
}

func parseBoard(s string) []int {
	result := []int{}
	for _, v := range strings.Split(s, " ") {
		if v != "" {
			c, _ := strconv.Atoi(v)
			result = append(result, c)
		}
	}
	return result
}
