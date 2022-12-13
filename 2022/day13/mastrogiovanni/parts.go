package main

import (
	"fmt"
	"sort"
)

func PartOne(inputs []string) string {
	tot := 0
	index := 1
	for i := 0; i < len(inputs); i += 3 {
		left := inputs[i]
		right := inputs[i+1]
		if res := Compare3(left, right); res {
			tot += index
		}
		index++
	}
	return fmt.Sprintf("%d", tot)
}

func PartTwo(inputs []string) string {
	list := make(Items, 0)
	for i := 0; i < len(inputs); i += 3 {
		left := parse(inputs[i])
		right := parse(inputs[i+1])
		list = append(list, left, right)
	}
	list = append(list, parse("[[2]]"), parse("[[6]]"))
	sort.Sort(list)
	tot := 1
	for i := 0; i < len(list); i++ {
		if fmt.Sprintf("%+v", list[i]) == "[[2]]" {
			tot *= (i + 1)
		}
		if fmt.Sprintf("%+v", list[i]) == "[[6]]" {
			tot *= (i + 1)
		}
	}
	return fmt.Sprintf("%d", tot)
}
