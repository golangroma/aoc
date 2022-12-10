package main

import (
	"sort"
	"strconv"
)

func PartOne(input []string) (int, error) {
	buff := 0
	arr := make([]int, 0)

	for _, v := range input {
		if v == "" {
			arr = append(arr, buff)
			buff = 0

			continue
		}

		number, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}

		buff += number
	}

	sort.Ints(arr)

	return arr[len(arr)-1], nil
}

func PartTwo(input []string) (int, error) {
	buff := 0
	arr := make([]int, 0)

	for _, v := range input {
		if v == "" {
			arr = append(arr, buff)
			buff = 0

			continue
		}

		number, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}

		buff += number
	}

	sort.Ints(arr)

	return arr[len(arr)-1] + arr[len(arr)-2] + arr[len(arr)-3], nil
}
