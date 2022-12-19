package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	StartingItems  []int
	Operation      func(int) int
	Div            int
	True           int
	False          int
	TotInspections int
}

func parseStartingItems(input string) []int {
	var items []int

	parsedInput := strings.Split(input, ": ")
	parsedItems := strings.Split(parsedInput[1], ", ")

	for _, item := range parsedItems {
		parsedItem, _ := strconv.Atoi(item)
		items = append(items, parsedItem)
	}

	return items
}

func parseOperation(input string) func(int) int {
	parsedInput := strings.Split(input, "= ")

	return func(old int) int {
		var parA, parB int

		parsedOp := strings.Split(parsedInput[1], " ")

		if parsedOp[0] == "old" {
			parA = old
		} else {
			parA, _ = strconv.Atoi(parsedOp[1])
		}

		if parsedOp[2] == "old" {
			parB = old
		} else {
			parB, _ = strconv.Atoi(parsedOp[2])
		}

		switch parsedOp[1] {
		case "+":
			return parA + parB
		case "*":
			return parA * parB
		case "-":
			return parA - parB
		case "/":
			return parA / parB
		}

		return 0
	}
}

func parseDiv(input string) int {
	parsedInput := strings.Split(input, ": ")
	parsedTest := strings.Split(parsedInput[1], " ")
	div, _ := strconv.Atoi(parsedTest[2])

	return div
}

func parseThrow(input string) int {
	parsedInput := strings.Split(input, ": ")
	parsedThrow := strings.Split(parsedInput[1], " ")
	monkeyIdx, _ := strconv.Atoi(parsedThrow[3])

	return monkeyIdx
}

func parseMonkeys(input []string) []Monkey {
	var monkeys []Monkey

	for i := 0; i < len(input); i += 7 {
		startingItems := parseStartingItems(input[i+1])
		operation := parseOperation(input[i+2])
		div := parseDiv(input[i+3])
		trueThrow := parseThrow(input[i+4])
		falseThrow := parseThrow(input[i+5])

		monkeys = append(monkeys, Monkey{
			StartingItems: startingItems,
			Operation:     operation,
			Div:           div,
			True:          trueThrow,
			False:         falseThrow,
		})
	}

	return monkeys
}

func monkeyRound(monkeys []Monkey, currentMonkey int, worryFunc func(int) int) {
	monkey := monkeys[currentMonkey]

	for i := 0; i < len(monkey.StartingItems); i++ {
		monkeys[currentMonkey].TotInspections++

		newWorryVal := monkey.Operation(monkey.StartingItems[i])
		monkey.StartingItems[i] = worryFunc(newWorryVal)

		if monkey.StartingItems[i]%monkey.Div == 0 {
			monkeys[monkey.True].StartingItems = append(monkeys[monkey.True].StartingItems, monkey.StartingItems[i])
			continue
		}

		monkeys[monkey.False].StartingItems = append(monkeys[monkey.False].StartingItems, monkey.StartingItems[i])
	}

	monkeys[currentMonkey].StartingItems = []int{}
}

func PartOne(input []string) string {
	monkeys := parseMonkeys(input)
	rounds := 20
	worryFunc := func(n int) int { return int(math.Floor(float64(n) / 3)) }

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeyRound(monkeys, j, worryFunc)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].TotInspections > monkeys[j].TotInspections
	})

	result := strconv.Itoa(monkeys[0].TotInspections * monkeys[1].TotInspections)

	return result
}

func PartTwo(input []string) string {
	monkeys := parseMonkeys(input)
	rounds := 10000

	m := 1
	for _, monkey := range monkeys {
		m *= monkey.Div
	}

	worryFunc := func(n int) int { return n % m }

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeyRound(monkeys, j, worryFunc)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].TotInspections > monkeys[j].TotInspections
	})

	result := strconv.Itoa(monkeys[0].TotInspections * monkeys[1].TotInspections)

	return result
}
