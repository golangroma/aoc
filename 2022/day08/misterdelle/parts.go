package main

import (
	"strconv"
)

func PartOne(input []string) string {
	counter := 0
	heightMap := loadHeightMap(input)

	for row := 0; row < len(heightMap); row++ {
		for col := 0; col < len(heightMap); col++ {
			if isVisibile(heightMap, row, col) {
				counter++
			}
		}
	}

	return strconv.Itoa(counter)
}

func PartTwo(input []string) string {
	result := 0
	heightMap := loadHeightMap(input)

	for row := 0; row < len(heightMap); row++ {
		for col := 0; col < len(heightMap); col++ {
			res := calculateScenicScore(heightMap, row, col)
			if res > result {
				result = res
			}
		}
	}

	return strconv.Itoa(result)
}

func loadHeightMap(input []string) map[int]map[int]int {
	var heightMap = make(map[int]map[int]int)

	for rowIndex, row := range input {
		numCols := len(row)
		heightMap[rowIndex] = make(map[int]int, numCols)
		for colIndex := 0; colIndex < numCols; colIndex++ {
			height, _ := strconv.Atoi(string(row[colIndex]))
			heightMap[rowIndex][colIndex] = height
		}
	}

	return heightMap
}

func isVisibile(heightMap map[int]map[int]int, row, col int) bool {
	if isVisibleEast(heightMap, row, col) || isVisibleNorth(heightMap, row, col) || isVisibleSouth(heightMap, row, col) || isVisibleWest(heightMap, row, col) {
		return true
	}
	return false
}

func isVisibleEast(heightMap map[int]map[int]int, row, col int) bool {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	height := heightMap[row][col]
	size := len(heightMap[row])

	for colIndex := col + 1; colIndex < size; colIndex++ {
		next := heightMap[row][colIndex]
		if next >= height {
			return false
		}
	}

	return true
}

func isVisibleWest(heightMap map[int]map[int]int, row, col int) bool {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	height := heightMap[row][col]

	for colIndex := col - 1; colIndex >= 0; colIndex-- {
		next := heightMap[row][colIndex]
		if next >= height {
			return false
		}
	}

	return true
}

func isVisibleNorth(heightMap map[int]map[int]int, row, col int) bool {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	height := heightMap[row][col]

	for rowIndex := row - 1; rowIndex >= 0; rowIndex-- {
		next := heightMap[rowIndex][col]
		if next >= height {
			return false
		}
	}

	return true
}

func isVisibleSouth(heightMap map[int]map[int]int, row, col int) bool {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	height := heightMap[row][col]
	size := len(heightMap[1])

	for rowIndex := row + 1; rowIndex < size; rowIndex++ {
		next := heightMap[rowIndex][col]
		if next >= height {
			return false
		}
	}

	return true
}

func calculateScenicScore(heightMap map[int]map[int]int, row, col int) int {
	//
	// Arriva questa matrice
	// 30373
	// 25512
	// 65332
	// 33549
	// 35390
	// con row 0 e col 0, 1, 2, 3...
	// con row 1 e col 0, 1, 2, 3...
	//

	scoreWest := 0
	scoreEast := 0
	scoreNorth := 0
	scoreSouth := 0

	if row == 0 || col == 0 || row == len(heightMap)-1 || row == len(heightMap)-1 {
		return 0
	}

	height := heightMap[row][col]

	//
	// North
	//
	for rowIndex := row - 1; rowIndex >= 0; rowIndex-- {
		next := heightMap[rowIndex][col]
		scoreNorth++
		if next >= height {
			break
		}
	}

	//
	// East
	//
	size := len(heightMap[row])

	for colIndex := col + 1; colIndex < size; colIndex++ {
		next := heightMap[row][colIndex]
		scoreEast++
		if next >= height {
			break
		}
	}

	//
	// South
	//
	size = len(heightMap[1])

	for rowIndex := row + 1; rowIndex < size; rowIndex++ {
		next := heightMap[rowIndex][col]
		scoreSouth++
		if next >= height {
			break
		}
	}

	//
	// West
	//
	for colIndex := col - 1; colIndex >= 0; colIndex-- {
		next := heightMap[row][colIndex]
		scoreWest++
		if next >= height {
			break
		}
	}

	return scoreNorth * scoreEast * scoreSouth * scoreWest
}
