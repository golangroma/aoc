package main

import (
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	cycleCounter := 1
	total := 0
	signalStrenght := 1
	value := 0
	cycleMap := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}

	for _, v := range input {
		value, _ = parseCommand(v)

		if value == 0 {
			cycleCounter++
			if cycleCounter == 20 || cycleCounter == 60 || cycleCounter == 100 || cycleCounter == 140 || cycleCounter == 180 || cycleCounter == 220 {
				cycleMap[cycleCounter] = signalStrenght * cycleCounter
			}
		} else {
			cycleCounter++
			if cycleCounter == 20 || cycleCounter == 60 || cycleCounter == 100 || cycleCounter == 140 || cycleCounter == 180 || cycleCounter == 220 {
				cycleMap[cycleCounter] = signalStrenght * cycleCounter
			}
			cycleCounter++
			signalStrenght += value
			if cycleCounter == 20 || cycleCounter == 60 || cycleCounter == 100 || cycleCounter == 140 || cycleCounter == 180 || cycleCounter == 220 {
				cycleMap[cycleCounter] = signalStrenght * cycleCounter
			}
		}
	}

	for _, v := range cycleMap {
		total += v
	}

	return strconv.Itoa(total)
}

func PartTwo(input []string) string {
	X := 1
	clock := 0
	result := ""
	value := 0
	length := 1
	for _, v := range input {
		value, length = parseCommand(v)

		//
		// Draws the CRT
		//
		for i := 0; i < length; i++ {
			currentCursor := clock % 40
			if currentCursor >= X-1 && currentCursor <= X+1 {
				result += "#"
			} else {
				result += "."
			}
			if i == length-1 {
				X = X + value
			}
			clock++
			if clock%40 < (clock-1)%40 {
				result += "\n"
			}
		}

	}

	return result
}

func parseCommand(input string) (int, int) {
	if input == "noop" {
		return 0, 1
	} else {
		commands := strings.Split(input, " ")
		value, _ := strconv.Atoi(commands[1])
		return value, 2
	}
}
