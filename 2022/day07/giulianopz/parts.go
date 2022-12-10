package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const (
	maxSize = 100000
)

type fs []*dir

type dir struct {
	size     int
	name     string
	children []int
	parent   int
}

func (fs *fs) insert(d *dir) int {

	dirIdx := fs.search(d.name)
	if dirIdx != -1 {
		([]*dir)(*fs)[dirIdx].size = d.size
		([]*dir)(*fs)[dirIdx].children = d.children
		return dirIdx
	}
	*fs = append(*fs, d)
	return len(*fs) - 1
}

func (fs *fs) propagate(fileSize, parentIdx int) {
	stack := make([]int, 0)

	if parentIdx != -1 {
		stack = append(stack, parentIdx)
	}

	for len(stack) != 0 {
		idx := stack[0]
		parent := ([]*dir)(*fs)[idx]
		parent.size += fileSize

		stack = stack[1:]
		if parent.parent != -1 {
			stack = append(stack, parent.parent)
		}
	}
}

func (fs *fs) search(name string) int {
	for i, d := range *fs {
		if d.name == name {
			return i
		}
	}
	return -1
}

const (
	cd = "cd"
	ls = "ls"

	dollar     = "$"
	fsRootName = "/"

	commandPattern = `^\$\s(\w+)\s?(.*)?`
	dirPattern     = `^dir\s(.*)$`
	filePattern    = `(\d+)\s(.*)`
)

func PartOne(input []string) string {

	commandRgx, err := regexp.Compile(commandPattern)
	if err != nil {
		log.Fatal(err)
	}

	dirRgx, err := regexp.Compile(dirPattern)
	if err != nil {
		log.Fatal(err)
	}

	fileRgx, err := regexp.Compile(filePattern)
	if err != nil {
		log.Fatal(err)
	}

	fs := fs{}
	var current *dir
	for i, line := range input {

		if commandRgx.MatchString(line) {
			groups := commandRgx.FindAllStringSubmatch(line, -1)
			cmd := groups[0][1]
			dirName := groups[0][2]
			if cmd == cd {

				if dirName != ".." {
					dirIdx := fs.search(dirName)
					if dirIdx != -1 {
						current = ([]*dir)(fs)[dirIdx]
					} else {
						current = new(dir)
						current.name = dirName
						current.parent = -1
						fs.insert(current)
					}
				}
			}
		} else if dirRgx.MatchString(line) {
			groups := dirRgx.FindAllStringSubmatch(line, -1)
			childName := groups[0][1]

			childIdx := fs.insert(&dir{name: childName, parent: fs.search(current.name)})
			current.children = append(current.children, childIdx)
		} else if fileRgx.MatchString(line) {
			groups := fileRgx.FindAllStringSubmatch(line, -1)
			fileSize, _ := strconv.Atoi(groups[0][1])
			current.size += fileSize
			fs.propagate(fileSize, current.parent)
		}

		if i == len(input)-1 {
			if current != nil {
				fs.insert(current)
			}
		}
	}

	var totSize int
	for _, d := range fs {
		if d.size <= maxSize {
			totSize += d.size
		}
	}

	return fmt.Sprintf("%d", totSize)
}

func PartTwo(input []string) string {
	return ""
}
