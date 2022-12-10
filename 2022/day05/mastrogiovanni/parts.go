package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(numbers []rune) []rune {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func ParseCargo(cargo []string) [][]rune {
	size := len(cargo[0])
	columns := (size + 1) / 4
	result := make([][]rune, columns)
	for j := range result {
		result[j] = make([]rune, 0)
	}
	for row := 0; row < len(cargo)-1; row++ {
		for i := 0; i < columns; i++ {
			l := cargo[row][i*4+1]
			if l != ' ' {
				result[i] = append(result[i], rune(l))
			}
		}
	}
	for j := 0; j < columns; j++ {
		result[j] = reverse(result[j])
	}
	return result
}

func ExtractCargo(input []string) ([]string, int) {
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			return input[0:i], i + 1
		}
	}
	return nil, 0
}

func dumpCargo(parsed [][]rune) {
	for j := range parsed {
		for i := range parsed[j] {
			fmt.Printf("%c", parsed[j][i])
		}
		fmt.Println()
	}
}

type Move struct {
	quantity int
	from     int
	to       int
}

func parseMove(line string) Move {
	line = strings.TrimSpace(line)
	components := strings.Split(line, " ")
	quantity, _ := strconv.Atoi(components[1])
	from, _ := strconv.Atoi(components[3])
	to, _ := strconv.Atoi(components[5])
	return Move{
		quantity: quantity,
		from:     from - 1,
		to:       to - 1,
	}
}

func applyMove(cargo [][]rune, move Move) {
	remain := cargo[move.from][:len(cargo[move.from])-move.quantity]
	to_remove := cargo[move.from][len(cargo[move.from])-move.quantity:]
	cargo[move.to] = append(cargo[move.to], reverse(to_remove)...)
	cargo[move.from] = remain
}

func applyMove9001(cargo [][]rune, move Move) {
	remain := cargo[move.from][:len(cargo[move.from])-move.quantity]
	to_remove := cargo[move.from][len(cargo[move.from])-move.quantity:]
	cargo[move.to] = append(cargo[move.to], to_remove...)
	cargo[move.from] = remain
}

func GetUpperCargo(cargo [][]rune) string {
	result := ""
	for j := 0; j < len(cargo); j++ {
		result = fmt.Sprintf("%s%c", result, cargo[j][len(cargo[j])-1])
	}
	return result
}

func PartOne(input []string) string {
	cargo, moves := ExtractCargo(input)
	parsed := ParseCargo(cargo)

	for i := moves; i < len(input); i++ {
		move := parseMove(input[i])
		applyMove(parsed, move)
	}

	return GetUpperCargo(parsed)
}

func PartTwo(input []string) string {
	cargo, moves := ExtractCargo(input)
	parsed := ParseCargo(cargo)

	for i := moves; i < len(input); i++ {
		move := parseMove(input[i])
		applyMove9001(parsed, move)
	}

	return GetUpperCargo(parsed)
}
