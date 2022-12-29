package main

import (
	"sort"
	"strconv"
	"strings"
)

type Operation struct {
	Expression string
}

type Monkey struct {
	Items           []int
	Op              Operation
	Divisor         int
	TrueIdx         int
	FalseIdx        int
	InspectionCount int
}

func PartOne(input []string) string {
	var monkeys []*Monkey

	monkeys = parseAllMonkeys(input, monkeys)

	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			monkeys = monkey.inspectItemsPartOne(monkeys)
		}
	}

	ic := getTwoMostInspected(monkeys)

	return strconv.Itoa(ic)
}

func PartTwo(input []string) string {
	var monkeys []*Monkey

	monkeys = parseAllMonkeys(input, monkeys)

	modulus := 1
	for _, monkey := range monkeys {
		modulus *= monkey.Divisor
	}

	for i := 0; i < 10_000; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			monkeys = monkey.inspectItemsPartTwo(monkeys, modulus)
		}
	}

	ic := getTwoMostInspected(monkeys)

	return strconv.Itoa(ic)
}

func NewOperation(expression string) *Operation {
	return &Operation{Expression: expression}
}

func NewMonkey(items ...int) *Monkey {
	var startingItems []int
	for i := 0; i < len(items); i++ {
		startingItems = append(startingItems, items[i])
	}

	return &Monkey{Items: startingItems}
}

func (o *Operation) applyPartOne(worryLevel int) int {
	//
	// old + 3
	// old * 3
	// old + old
	// old * old
	//
	tokens := strings.Split(o.Expression, " ")
	operator := tokens[1]
	right := tokens[2]
	rightVal := worryLevel

	if operator == "+" {
		// Somma
		if right != "old" {
			rightVal, _ = strconv.Atoi(right)
		}
		return (worryLevel + rightVal) / 3
	} else if operator == "*" {
		// Moltiplicazione
		if right != "old" {
			rightVal, _ = strconv.Atoi(right)
		}
		return (worryLevel * rightVal) / 3
	}

	return 0
}

func (o *Operation) applyPartTwo(worryLevel, m int) int {
	//
	// old + 3
	// old * 3
	// old + old
	// old * old
	//
	tokens := strings.Split(o.Expression, " ")
	operator := tokens[1]
	right := tokens[2]
	rightVal := worryLevel

	if operator == "+" {
		// Somma
		if right != "old" {
			rightVal, _ = strconv.Atoi(right)
		}
		return (worryLevel + rightVal) % m
	} else if operator == "*" {
		// Moltiplicazione
		if right != "old" {
			rightVal, _ = strconv.Atoi(right)
		}
		return (worryLevel * rightVal) % m
	}

	return 0
}

func (m *Monkey) inspectItemsPartOne(friends []*Monkey) []*Monkey {
	for i := 0; i < len(m.Items); i++ {
		item := m.Items[i]

		newWorryLevel := m.Op.applyPartOne(item)

		if newWorryLevel%m.Divisor == 0 {
			friends[m.TrueIdx].Items = append(friends[m.TrueIdx].Items, newWorryLevel)
		} else {
			friends[m.FalseIdx].Items = append(friends[m.FalseIdx].Items, newWorryLevel)
		}
		m.InspectionCount++
	}

	m.Items = nil

	return friends
}

func (m *Monkey) inspectItemsPartTwo(friends []*Monkey, modulus int) []*Monkey {
	for i := 0; i < len(m.Items); i++ {
		item := m.Items[i]

		newWorryLevel := m.Op.applyPartTwo(item, modulus)

		if newWorryLevel%m.Divisor == 0 {
			friends[m.TrueIdx].Items = append(friends[m.TrueIdx].Items, newWorryLevel)
		} else {
			friends[m.FalseIdx].Items = append(friends[m.FalseIdx].Items, newWorryLevel)
		}
		m.InspectionCount++
	}

	m.Items = nil

	return friends
}

func parseItems(input string) []int {
	// Starting items: 79, 98
	var parsedItems []int
	if strings.HasPrefix(input, "Starting items: ") {
		tokens := strings.Split(input, ":")
		items := strings.Split(tokens[1], ",")
		for i := 0; i < len(items); i++ {
			itemInt, _ := strconv.Atoi(strings.Trim(items[i], " "))
			parsedItems = append(parsedItems, itemInt)
		}
	}

	return parsedItems
}

func parseOperation(input string) string {
	// Operation: new = old * 19
	if strings.HasPrefix(input, "Operation: ") {
		tokens := strings.Split(input, "=")
		return strings.Trim(tokens[1], " ")
	}

	return ""
}

func parseTest(input string) int {
	// Test: divisible by 23
	var parsedTest int = 0
	if strings.HasPrefix(input, "Test: ") {
		tokens := strings.Split(input, " ")
		parsedTest, _ = strconv.Atoi(strings.Trim(tokens[3], " "))
	}

	return parsedTest
}

func parseTrueIdx(input string) int {
	// If true: throw to monkey 2
	var parsedIdx int = 0
	if strings.HasPrefix(input, "If true: ") {
		tokens := strings.Split(input, " ")
		parsedIdx, _ = strconv.Atoi(strings.Trim(tokens[len(tokens)-1], " "))
	}

	return parsedIdx
}

func parseFalseIdx(input string) int {
	// If false: throw to monkey 3
	var parsedIdx int = 0
	if strings.HasPrefix(input, "If false: ") {
		tokens := strings.Split(input, " ")
		parsedIdx, _ = strconv.Atoi(strings.Trim(tokens[len(tokens)-1], " "))
	}

	return parsedIdx
}

func parseAllMonkeys(input []string, monkeys []*Monkey) []*Monkey {
	// Monkey 0:
	// 	Starting items: 79, 98
	// 	Operation: new = old * 19
	// 	Test: divisible by 23
	//  	If true: throw to monkey 2
	//  	If false: throw to monkey 3

	var monkey *Monkey

	for _, v := range input {
		vv := strings.Trim(v, " ")
		if strings.HasPrefix(vv, "Monkey ") {
			// Inizio Monkey
			monkey = NewMonkey()
			monkeys = append(monkeys, monkey)
		} else {
			if len(parseItems(vv)) > 0 {
				monkey.Items = parseItems(vv)
			}

			if parseOperation(vv) != "" {
				monkey.Op.Expression = parseOperation(vv)
			}

			if parseTest(vv) > 0 {
				monkey.Divisor = parseTest(vv)
			}

			if parseTrueIdx(vv) > 0 {
				monkey.TrueIdx = parseTrueIdx(vv)
			}

			if parseFalseIdx(vv) > 0 {
				monkey.FalseIdx = parseFalseIdx(vv)
			}
		}
	}

	return monkeys
}

func getTwoMostInspected(friends []*Monkey) int {
	var inspections []int

	for i := 0; i < len(friends); i++ {
		inspections = append(inspections, friends[i].InspectionCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))

	return inspections[0] * inspections[1]
}
