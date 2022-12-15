package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	totalRangeFullyContained := 0
	count := 0

	for _, v := range input {
		pairs := strings.Split(v, ",")
		left := pairs[0]
		right := pairs[1]
		pairsLeft := strings.Split(left, "-")
		pairsRight := strings.Split(right, "-")

		startLeft, _ := strconv.Atoi(pairsLeft[0])
		endLeft, _ := strconv.Atoi(pairsLeft[1])
		startRight, _ := strconv.Atoi(pairsRight[0])
		endRight, _ := strconv.Atoi(pairsRight[1])

		//
		// 2-8,3-7
		// 2-8 fully contains 3-7
		//
		// 6-6,4-6
		// 6-6 is fully contained by 4-6
		//
		// 15-57,15-57
		// 15-57 is fully contained by 15-57
		//

		if startRight == startLeft && endRight == endLeft {
			//
			// If the ranges are exactly the same I count as one
			//
			totalRangeFullyContained++
		} else {
			if startRight >= startLeft {
				if endRight <= endLeft {
					//
					// Right pair completely contained in Left pair
					//
					totalRangeFullyContained++
				}
			}

			if startLeft >= startRight {
				if endLeft <= endRight {
					//
					// Left pair completely contained in Right pair
					//
					totalRangeFullyContained++
				}
			}
		}

		if (startRight-startLeft)*(endLeft-endRight) >= 0 {
			count++
		}
	}

	fmt.Println(count)
	fmt.Println(totalRangeFullyContained)

	return strconv.Itoa(count)
}

func PartTwo(input []string) string {
	var startLeft, endLeft, startRight, endRight int
	totalRangeFullyContained := 0

	for _, v := range input {
		fmt.Sscanf(v, "%d-%d,%d-%d\n", &startLeft, &endLeft, &startRight, &endRight)

		if (endRight-startLeft)*(endLeft-startRight) >= 0 {
			totalRangeFullyContained++
		}
	}

	return strconv.Itoa(totalRangeFullyContained)
}
