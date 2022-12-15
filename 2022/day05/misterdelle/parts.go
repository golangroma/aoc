package main

import (
	"fmt"
)

type Stack struct {
	boxes []rune
}

func (c *Stack) Push(v rune) {
	c.boxes = append(c.boxes, v)
}

func (c *Stack) Pop() rune {
	newStack := c.boxes[len(c.boxes)-1]
	c.boxes = c.boxes[:len(c.boxes)-1]

	return newStack
}

func (c *Stack) PushN(v []rune) {
	c.boxes = append(c.boxes, v...)
}

func (c *Stack) PopN(x int) []rune {
	newStack := c.boxes[len(c.boxes)-x : len(c.boxes)]
	c.boxes = c.boxes[:len(c.boxes)-x]

	return newStack
}

func PartOne(input []string) string {
	createsStartingIndex := getMovementStartingIndex(input)
	stacks := initStacks(input[:createsStartingIndex])

	for _, v := range input[createsStartingIndex+1:] {
		//
		// Parse movement indications
		//
		var nItemsToMove, fromCrane, toCrane int
		fmt.Sscanf(v, "move %d from %d to %d", &nItemsToMove, &fromCrane, &toCrane)

		for i := 0; i < nItemsToMove; i++ {
			stacks[toCrane].Push(stacks[fromCrane].Pop())
		}
	}

	return getFirstElementOfEachStack(stacks)
}

func PartTwo(input []string) string {
	movementStartingIndex := getMovementStartingIndex(input)
	stacks := initStacks(input[:movementStartingIndex])

	for _, v := range input[movementStartingIndex+1:] {
		//
		// Parse movement indications
		//
		var nItemsToMove, fromCrane, toCrane int
		fmt.Sscanf(v, "move %d from %d to %d", &nItemsToMove, &fromCrane, &toCrane)

		stacks[toCrane].PushN(stacks[fromCrane].PopN(nItemsToMove))
	}

	return getFirstElementOfEachStack(stacks)
}

func getMovementStartingIndex(input []string) int {
	for i, v := range input {
		if v == "" {
			return i
		}
	}

	return 0
}

func initStacks(input []string) map[int]*Stack {
	stacks := make(map[int]*Stack)
	for i := 1; i < len(input[len(input)-1]); i += 4 {
		stacks[(i/4)+1] = &Stack{}
	}

	for i := len(input) - 1; i >= 0; i-- {
		for j := 1; j < len(input[i]); j += 4 {
			item := rune(input[i][j])
			if item != ' ' {
				stacks[(j/4)+1].Push(item)
			}
		}
	}

	return stacks
}

func getFirstElementOfEachStack(stacks map[int]*Stack) string {
	var ret string = ""
	for i := 1; i <= len(stacks); i++ {
		ret += string(stacks[i].Pop())
	}

	return ret
}
