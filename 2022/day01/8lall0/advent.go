package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type Elf struct {
	Num      int
	Calories int
}

func (e Elf) String() string {
	return fmt.Sprintf("Elf n° %d with %d calories", e.Num, e.Calories)
}

func part1(r io.Reader) (Elf, error) {
	var curElf Elf

	scanner := bufio.NewScanner(r)
	nElf := 1
	sum := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			if curElf.Calories < sum {
				curElf.Num = nElf
				curElf.Calories = sum
			}
			nElf++
			sum = 0
		} else {
			val, err := strconv.Atoi(txt)
			if err != nil {
				return Elf{}, err
			}

			sum += val
		}
	}

	if curElf.Calories < sum {
		curElf.Num = nElf
		curElf.Calories = sum
	}

	return curElf, nil
}

func part2(r io.Reader) ([]Elf, error) {
	elves := make([]Elf, 0)

	scanner := bufio.NewScanner(r)
	nElf := 1
	sum := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			elves = append(elves, Elf{Num: nElf, Calories: sum})
			nElf++
			sum = 0
		} else {
			val, err := strconv.Atoi(txt)
			if err != nil {
				return nil, err
			}

			sum += val
		}
	}

	elves = append(elves, Elf{Num: nElf, Calories: sum})

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	return elves[0:3], nil
}
