package main

import (
	"strconv"
	"strings"
)

//
// Overall Strategy
//
// 1) Create ElfFile Data Model
// 2) Create ElfDirectory Data Model
// 3) Parse input into Data Model
// 4) Write a function to get all ElfDirectories  in our model
// 5) Iterate over each ElfDirectory:
//    - Sum all directories that meet criteria
//

const (
	maxDirSize      = 100_000
	diskSize        = 70_000_000
	neededFreeSpace = 30_000_000
)

func PartOne(input []string) string {
	fs := buildFileSystem(input)
	result := getLimitedTotalDirectorySize(&fs)
	return strconv.Itoa(result)
}

func PartTwo(input []string) string {
	fs := buildFileSystem(input)
	diskOccupation := fs.size()
	toFree := diskOccupation - (diskSize - neededFreeSpace)
	result := findSmallestDirectoryBySize(&fs, diskOccupation, toFree)
	return strconv.Itoa(result)
}

func buildFileSystem(input []string) ElfDirectory {
	var root ElfDirectory = newElfDirectory("/", nil)
	var currDir = &root
	ef := ElfFile{}
	ed := ElfDirectory{}

	for _, row := range input {
		if ef.isParsable(row) {
			processFileLine(currDir, row)
		} else if ed.isParsable(row) {
			processDirectoryLine(currDir, row)
		} else if strings.HasPrefix(row, "$ cd") {
			currDir = processChangeDirectory(currDir, row)
		}
	}

	return root
}

func processDirectoryLine(currDir *ElfDirectory, input string) {
	parseElfDirectory(input, currDir)
}

func processFileLine(currDir *ElfDirectory, input string) {
	ef := parseElfFile(input)
	currDir.addFile(ef)
}

func processChangeDirectory(currDir *ElfDirectory, input string) *ElfDirectory {
	childName := strings.Split(input, " ")[2]

	if childName == ".." {
		if currDir.parent != nil {
			return currDir.parent
		} else {
			return currDir
		}
	} else if childName == "/" {
		return currDir
	} else {
		v, ok := currDir.children[childName]
		if ok {
			return v
		} else {
			return &ElfDirectory{}
		}
	}
}

func getLimitedTotalDirectorySize(ed *ElfDirectory) int {
	dirs := ed.findAllDirectories()

	size := 0
	for _, v := range dirs {
		if v.size() < maxDirSize {
			size += v.size()
		}
	}

	return size
}

func findSmallestDirectoryBySize(node *ElfDirectory, best int, freeSpace int) int {
	for _, child := range node.children {
		if child.size() >= freeSpace {
			if child.size() < best {
				best = child.size()
			}
		}
		best = findSmallestDirectoryBySize(child, best, freeSpace)
	}
	return best
}
