package main

import (
	"log"
	"regexp"
	"strconv"
)

type fs struct {
	root *dir
}

type dir struct {
	size     int
	name     string
	children []*dir
	parent   *dir
}

func (fs *fs) insert(d *dir) {

	if fs.root == nil {
		fs.root = d
		return
	}

	stack := make([]*dir, 0)
	stack = append(stack, fs.root.children...)
	for len(stack) != 0 {
		current := stack[0]
		if current.name == d.name {
			//TODO propagate size
			current.size = d.size
			updateSize(current)
			current.children = d.children
			return
		}
		stack = append(stack, current.children...)
		stack = stack[1:]
	}
}

func updateSize(d *dir) {
	size := d.size
	stack := make([]*dir, 0)
	stack = append(stack, d.parent)
	for len(stack) != 0 {
		current := stack[0]
		current.size += size
		size = current.size

		stack = stack[1:]
		if current.parent != nil {
			stack = append(stack, current.parent)
		}
	}
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
	for _, line := range input {

		if commandRgx.MatchString(line) {
			groups := commandRgx.FindAllStringSubmatch(line, -1)
			cmd := groups[0][1]
			dirName := groups[0][2]
			if cmd == cd {

				if dirName == ".." {
					continue
				}

				if dirName == "/" {
					current = new(dir)
					current.name = dirName
				} else {
					fs.insert(current)
					current = new(dir)
					current.name = dirName
				}
			}
		} else if dirRgx.MatchString(line) {
			groups := dirRgx.FindAllStringSubmatch(line, -1)
			childName := groups[0][1]
			current.children = append(current.children, &dir{name: childName, parent: current})
		} else if fileRgx.MatchString(line) {
			groups := fileRgx.FindAllStringSubmatch(line, -1)
			fileSize, _ := strconv.Atoi(groups[0][1])
			current.size += fileSize
		}

	}
	return ""
}

func PartTwo(input []string) string {
	return ""
}
