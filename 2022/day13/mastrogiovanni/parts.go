package main

import (
	"fmt"
)

func PartOne(inputs []string) string {
	tot := 0
	index := 1
	for i := 0; i < len(inputs); i += 3 {
		left := inputs[i]
		right := inputs[i+1]
		fmt.Println(left)
		fmt.Println(right)
		fmt.Printf("\n\n%s\n%s\n", left, right)
		// res := Compare1(left, right)
		// res := Compare2(left, right)
		res := Compare3(left, right)
		fmt.Printf("%+v\n\n", res)
		if res {
			tot += index
		}
		index++
	}
	return fmt.Sprintf("%d", tot)
}

func PartTwo(input []string) string {
	return ""
}
