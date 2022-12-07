package main

import (
	"strconv"
	"strings"
)

type Dir struct {
	Files   []File
	SubDirs []Dir
	Parent  *Dir
	Name    string
	Size    int
}

type File struct {
	Name string
	Size int
}

func PartOne(input []string) string {
	var result int

	root := buildTree(input)

	fillSizes(root)

	getSizesSumAtMost(root, 100000, &result)

	return strconv.Itoa(result)
}

func buildTree(input []string) *Dir {
	var root *Dir
	var currentDir *Dir

	rootDir := Dir{Name: "/"}

	root = &rootDir
	currentDir = root

	for _, line := range input[2:] {
		if line[0] == '$' {
			cmd := strings.Split(line, " ")

			switch cmd[1] {
			case "cd":
				if cmd[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for j, dir := range currentDir.SubDirs {
						if dir.Name == cmd[2] {
							currentDir = &currentDir.SubDirs[j]
							break
						}
					}
				}

				continue
			default:
				continue
			}
		}

		vals := strings.Split(line, " ")

		if vals[0] == "dir" {
			currentDir.SubDirs = append(currentDir.SubDirs, Dir{Name: vals[1], Parent: currentDir})
			continue
		}

		size, _ := strconv.Atoi(vals[0])

		currentDir.Files = append(currentDir.Files, File{Name: vals[1], Size: size})
	}

	return root
}

func fillSizes(dir *Dir) {
	for _, file := range dir.Files {
		dir.Size += file.Size
	}

	for i := range dir.SubDirs {
		fillSizes(&dir.SubDirs[i])
		dir.Size += dir.SubDirs[i].Size
	}
}

func getSizesSumAtMost(dir *Dir, max int, sum *int) {
	for i, subDir := range dir.SubDirs {
		if subDir.Size <= max {
			*sum += subDir.Size
		}

		getSizesSumAtMost(&dir.SubDirs[i], max, sum)
	}
}

func PartTwo(input []string) string {
	var result int

	totalDiskSize := 70000000
	wantUnused := 30000000

	root := buildTree(input)

	fillSizes(root)

	actualFreeDisk := totalDiskSize - root.Size
	atLeast := wantUnused - actualFreeDisk

	result = root.Size

	getMinUnusedAtLeast(root, atLeast, &result)

	return strconv.Itoa(result)
}

func getMinUnusedAtLeast(dir *Dir, atLeast int, min *int) {

	for i, subDir := range dir.SubDirs {
		if subDir.Size >= atLeast && subDir.Size < *min {
			*min = subDir.Size
		}

		getMinUnusedAtLeast(&dir.SubDirs[i], atLeast, min)
	}
}
