package main

import (
	"fmt"
	"sort"
)

type MonkeyOperation func(old int) int

type Monkey struct {
	worry_levels    []int
	operation       MonkeyOperation
	divisor         int
	monkey_if_true  int
	monkey_if_false int
	inspection      int
}

func DumpMonkeys(monkeys []*Monkey) {
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d {insp. %d}: %+v\n", i, monkeys[i].inspection, monkeys[i].worry_levels)
	}
}

func MonkeyProcessor(monkeys []*Monkey) {
	for index := 0; index < len(monkeys); index++ {
		current := monkeys[index]
		for _, wl := range current.worry_levels {
			newWl := current.operation(wl)
			newWl = newWl / 3
			if newWl%current.divisor == 0 {
				monkeys[current.monkey_if_true].worry_levels = append(monkeys[current.monkey_if_true].worry_levels, newWl)
			} else {
				monkeys[current.monkey_if_false].worry_levels = append(monkeys[current.monkey_if_false].worry_levels, newWl)
			}
			current.inspection++
		}
		current.worry_levels = []int{}
	}
}

func MonkeyProcessor2(big_factor int, monkeys []*Monkey) {
	for index := 0; index < len(monkeys); index++ {
		current := monkeys[index]
		for _, wl := range current.worry_levels {
			current.inspection++
			newWl := current.operation(wl)
			newWl = newWl % big_factor
			if newWl%current.divisor == 0 {
				monkeys[current.monkey_if_true].worry_levels = append(monkeys[current.monkey_if_true].worry_levels, newWl)
			} else {
				monkeys[current.monkey_if_false].worry_levels = append(monkeys[current.monkey_if_false].worry_levels, newWl)
			}
		}
		current.worry_levels = []int{}
	}
}

func GetMonkeyBusiness(monkeys []*Monkey) int {
	res := make([]int, len(monkeys))
	for i, m := range monkeys {
		res[i] = m.inspection
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res)))
	return res[0] * res[1]
}
