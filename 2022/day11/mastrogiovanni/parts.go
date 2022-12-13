package main

import (
	"fmt"
)

func PartOne(monkeys []*Monkey) string {
	for i := 0; i < 20; i++ {
		MonkeyProcessor(monkeys)
	}
	return fmt.Sprintf("%d", GetMonkeyBusiness(monkeys))
}

func PartTwo(monkeys []*Monkey) string {
	big_factor := 1
	for index := 0; index < len(monkeys); index++ {
		big_factor *= monkeys[index].divisor
	}
	for i := 0; i < 10000; i++ {
		MonkeyProcessor2(big_factor, monkeys)
	}
	return fmt.Sprintf("%d", GetMonkeyBusiness(monkeys))
}
