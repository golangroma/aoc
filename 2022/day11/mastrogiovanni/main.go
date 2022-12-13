package main

import "fmt"

func parse() []*Monkey {
	m0 := &Monkey{
		worry_levels:    []int{73, 77},
		operation:       func(old int) int { return old * 5 },
		divisor:         11,
		monkey_if_true:  6,
		monkey_if_false: 5,
	}

	m1 := &Monkey{
		worry_levels:    []int{57, 88, 80},
		operation:       func(old int) int { return old + 5 },
		divisor:         19,
		monkey_if_true:  6,
		monkey_if_false: 0,
	}

	m2 := &Monkey{
		worry_levels:    []int{61, 81, 84, 69, 77, 88},
		operation:       func(old int) int { return old * 19 },
		divisor:         5,
		monkey_if_true:  3,
		monkey_if_false: 1,
	}

	m3 := &Monkey{
		worry_levels:    []int{78, 89, 71, 60, 81, 84, 87, 75},
		operation:       func(old int) int { return old + 7 },
		divisor:         3,
		monkey_if_true:  1,
		monkey_if_false: 0,
	}

	m4 := &Monkey{
		worry_levels:    []int{60, 76, 90, 63, 86, 87, 89},
		operation:       func(old int) int { return old + 2 },
		divisor:         13,
		monkey_if_true:  2,
		monkey_if_false: 7,
	}

	m5 := &Monkey{
		worry_levels:    []int{88},
		operation:       func(old int) int { return old + 1 },
		divisor:         17,
		monkey_if_true:  4,
		monkey_if_false: 7,
	}

	m6 := &Monkey{
		worry_levels:    []int{84, 98, 78, 85},
		operation:       func(old int) int { return old * old },
		divisor:         7,
		monkey_if_true:  5,
		monkey_if_false: 4,
	}

	m7 := &Monkey{
		worry_levels:    []int{98, 89, 78, 73, 71},
		operation:       func(old int) int { return old + 4 },
		divisor:         2,
		monkey_if_true:  3,
		monkey_if_false: 2,
	}

	return []*Monkey{m0, m1, m2, m3, m4, m5, m6, m7}
}

func main() {
	monkeys := parse()
	fmt.Printf("Part 1: %v\n", PartOne(monkeys))

	monkeys = parse()
	fmt.Printf("Part 1: %v\n", PartTwo(monkeys))
}
