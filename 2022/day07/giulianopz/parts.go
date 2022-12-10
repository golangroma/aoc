package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

const (
	maxSize = 100000

	cd = "cd"
	ls = "ls"

	dollarSign  = "$"
	fsRootName  = "/"
	previousDir = ".."

	commandPattern = `^\$\s(\w+)\s?(.*)?`
	dirPattern     = `^dir\s(.*)$`
	filePattern    = `(\d+)\s(.*)`
)

type dir struct {
	size     int
	name     string
	children []*dir
	parent   *dir
}

func propagate(fileSize int, parent *dir) {
	stack := make([]*dir, 0)

	if parent != nil {
		stack = append(stack, parent)
	}

	for len(stack) != 0 {
		p := stack[0]
		p.size += fileSize
		stack = stack[1:]

		if p.parent != nil {
			stack = append(stack, p.parent)
		}
	}
}

func regex(patter string) *regexp.Regexp {
	regex, err := regexp.Compile(patter)
	if err != nil {
		log.Fatal(err)
	}
	return regex
}

func PartOne(input []string) string {

	commandRgx := regex(commandPattern)
	dirRgx := regex(dirPattern)
	fileRgx := regex(filePattern)

	var root *dir

	var lastDir *dir
	for _, line := range input {

		if commandRgx.MatchString(line) {
			groups := commandRgx.FindAllStringSubmatch(line, -1)
			cmd := groups[0][1]
			dirName := groups[0][2]

			if cmd == ls {
				continue
			}
			if cmd == cd {
				if lastDir == nil {
					lastDir = new(dir)
					lastDir.name = dirName
					if dirName == fsRootName {
						root = lastDir
					}
				} else {
					if dirName == previousDir {
						lastDir = lastDir.parent
					}
					for _, c := range lastDir.children {
						if c.name == dirName {
							child := c
							lastDir = child
							break
						}
					}
				}
			}

		} else if dirRgx.MatchString(line) {
			groups := dirRgx.FindAllStringSubmatch(line, -1)
			childName := groups[0][1]

			child := new(dir)
			child.name = childName
			child.parent = lastDir
			lastDir.children = append(lastDir.children, child)

		} else if fileRgx.MatchString(line) {
			groups := fileRgx.FindAllStringSubmatch(line, -1)
			fileSize, _ := strconv.Atoi(groups[0][1])
			lastDir.size += fileSize
			propagate(fileSize, lastDir.parent)
		}
	}

	var totSize int
	stack := make([]*dir, 0)
	stack = append(stack, root)
	for len(stack) != 0 {

		d := stack[0]
		if d.size <= maxSize {
			totSize += d.size
		}
		stack = stack[1:]

		if len(d.children) != 0 {
			stack = append(stack, d.children...)
		}
	}

	return fmt.Sprintf("%d", totSize)
}

const (
	maxDiskSpace = 70000000
	updateSize   = 30000000
)

func PartTwo(input []string) string {
	commandRgx := regex(commandPattern)
	dirRgx := regex(dirPattern)
	fileRgx := regex(filePattern)

	var root *dir

	var lastDir *dir
	for _, line := range input {

		if commandRgx.MatchString(line) {
			groups := commandRgx.FindAllStringSubmatch(line, -1)
			cmd := groups[0][1]
			dirName := groups[0][2]

			if cmd == ls {
				continue
			}
			if cmd == cd {
				if lastDir == nil {
					lastDir = new(dir)
					lastDir.name = dirName
					if dirName == fsRootName {
						root = lastDir
					}
				} else {
					if dirName == previousDir {
						lastDir = lastDir.parent
					}
					for _, c := range lastDir.children {
						if c.name == dirName {
							child := c
							lastDir = child
							break
						}
					}
				}
			}

		} else if dirRgx.MatchString(line) {
			groups := dirRgx.FindAllStringSubmatch(line, -1)
			childName := groups[0][1]

			child := new(dir)
			child.name = childName
			child.parent = lastDir
			lastDir.children = append(lastDir.children, child)

		} else if fileRgx.MatchString(line) {
			groups := fileRgx.FindAllStringSubmatch(line, -1)
			fileSize, _ := strconv.Atoi(groups[0][1])
			lastDir.size += fileSize
			propagate(fileSize, lastDir.parent)
		}
	}

	var sizes []int
	stack := make([]*dir, 0)
	stack = append(stack, root)
	for len(stack) != 0 {

		d := stack[0]
		sizes = append(sizes, d.size)
		stack = stack[1:]

		if len(d.children) != 0 {
			stack = append(stack, d.children...)
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	freeSpace := maxDiskSpace - root.size
	for _, s := range sizes {
		if freeSpace+s >= updateSize {
			return fmt.Sprintf("%d", s)
		}
	}
	return fmt.Sprintf("%d", 0)
}
