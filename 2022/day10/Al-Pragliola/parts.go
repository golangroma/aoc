package main

import (
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	result := 0
	regX := 1
	opIdx := 0
	workingOp := 0

	opDurations := map[string]int{
		"addx": 1,
		"noop": 0,
	}

	for cycles := 0; cycles < 220; cycles++ {
		switch cycles + 1 {
		case 20:
			result += (cycles + 1) * regX
		case 60:
			result += (cycles + 1) * regX
		case 100:
			result += (cycles + 1) * regX
		case 140:
			result += (cycles + 1) * regX
		case 180:
			result += (cycles + 1) * regX
		case 220:
			result += (cycles + 1) * regX
		}

		if opIdx == len(input) {
			continue
		}

		op := input[opIdx]
		opParts := strings.Split(op, " ")

		if workingOp < opDurations[opParts[0]] {
			workingOp++
			continue
		}

		if workingOp == opDurations[opParts[0]] {
			if opParts[0] == "addx" {
				val, _ := strconv.Atoi(opParts[1])

				regX += val
			}

			workingOp = 0
			opIdx++
			continue
		}

	}

	return strconv.Itoa(result)
}

func PartTwo(input []string) string {
	result := ""
	regX := 1
	opIdx := 0
	workingOp := 0

	opDurations := map[string]int{
		"addx": 1,
		"noop": 0,
	}

	crt := make(map[int]byte, 241)

	for cycles := 0; cycles < 240; cycles++ {
		currentPixel := '.'

		if cycles%40 >= (regX-1)%40 && cycles%40 <= (regX+1)%40 {
			currentPixel = '#'
		}

		crt[cycles] = byte(currentPixel)

		if opIdx == len(input) {
			continue
		}

		op := input[opIdx]
		opParts := strings.Split(op, " ")

		if workingOp < opDurations[opParts[0]] {
			workingOp++
			continue
		}

		if workingOp == opDurations[opParts[0]] {
			if opParts[0] == "addx" {
				val, _ := strconv.Atoi(opParts[1])

				regX += val
			}

			workingOp = 0
			opIdx++
			continue
		}
	}

	for i := 0; i < 240; i++ {
		result += string(crt[i])
		if (i+1)%40 == 0 {
			result += "\n"
		}
	}

	return result
}
