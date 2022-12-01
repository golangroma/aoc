package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/utils"
)

type Elf struct {
	ID    int
	Foods []int
}

func NewElf(id int) *Elf {
	return &Elf{
		ID:    id,
		Foods: []int{},
	}
}

func (e *Elf) Calories() int {
	calories := 0
	for i := range e.Foods {
		calories += e.Foods[i]
	}
	return calories
}

func (e *Elf) AddFood(calories int) {
	e.Foods = append(e.Foods, calories)
}

func (e *Elf) String() string {
	foods := []string{}
	for _, food := range e.Foods {
		foods = append(foods, strconv.Itoa(food))
	}
	return fmt.Sprintf("Elf{ID:%d, Foods:[%s]}", e.ID, strings.Join(foods, ","))
}

func PartOne(input []string) string {
	elves := elvesFromInput(input)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories() > elves[j].Calories()
	})

	return strconv.Itoa(elves[0].Calories())
}

func PartTwo(input []string) string {
	elves := elvesFromInput(input)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories() > elves[j].Calories()
	})

	topThreeTotal := elves[0].Calories() + elves[1].Calories() + elves[2].Calories()

	return strconv.Itoa(topThreeTotal)
}

func elvesFromInput(input []string) []*Elf {
	elves := []*Elf{}

	var elf *Elf
	for _, caloriesStr := range input {
		if elf == nil {
			// the ID of the Elf is derived from the order/length of the slice
			elf = NewElf(len(elves) + 1)
		}

		if caloriesStr == "" {
			elves = append(elves, elf)
			elf = nil
			continue
		}

		calories, err := strconv.Atoi(caloriesStr)
		utils.CheckErr(err)
		elf.AddFood(calories)
	}
	elves = append(elves, elf)

	return elves
}
