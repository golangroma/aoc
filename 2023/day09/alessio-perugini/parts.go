package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(input []string) [][]int {
	r := make([][]int, len(input))
	for i, v := range input {
		s := strings.Split(v, " ")
		elms := make([]int, len(s))
		for j, e := range s {
			elms[j] = mustConvertToInt(e)
		}
		r[i] = elms
	}
	return r
}

func PartOne(input []string) string {
	p := Parse(input)
	finalResult := 0
	for _, v := range p {
		r := sub([][]int{}, false, v)
		rightDiag := []int{0}
		for i := len(r) - 2; i >= 0; i-- {
			rightDiag = append(rightDiag, r[i][len(r[i])-1]+rightDiag[len(rightDiag)-1])
		}
		rightDiag = append(rightDiag, rightDiag[len(rightDiag)-1]+v[len(v)-1])
		finalResult += rightDiag[len(rightDiag)-1]
	}
	return fmt.Sprintf("%d", finalResult)
}

func sub(result [][]int, allZero bool, input []int) [][]int {
	if allZero {
		return result
	}
	subResult := []int{}
	allZero = true
	for i := 0; i < len(input)-1; i++ {
		diff := input[i+1] - input[i]
		subResult = append(subResult, diff)
		if diff != 0 {
			allZero = false
		}
	}
	result = append(result, subResult)
	return sub(result, allZero, subResult)
}

func PartTwo(input []string) string {
	p := Parse(input)
	finalResult := 0
	for _, v := range p {
		r := sub([][]int{}, false, v)
		leftDiag := []int{0}
		for i := len(r) - 2; i >= 0; i-- {
			leftDiag = append(leftDiag, r[i][0]-leftDiag[len(leftDiag)-1])
		}
		leftDiag = append(leftDiag, v[0]-leftDiag[len(leftDiag)-1])
		finalResult += leftDiag[len(leftDiag)-1]
	}
	return fmt.Sprintf("%d", finalResult)
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
