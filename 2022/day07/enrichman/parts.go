package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/utils"
)

type FileType int

const (
	Directory FileType = iota + 1
	RegularFile
)

type File struct {
	Name   string
	Type   FileType
	Size   int
	File   []*File
	Parent *File
}

func NewFile(name string, _type FileType, size int) *File {
	return &File{
		Name: name,
		Type: _type,
		Size: size,
		File: []*File{},
	}
}

func (f *File) CD(dir string) *File {
	if dir == ".." {
		return f.Parent
	}

	for _, f := range f.File {
		if f.Name == dir {
			return f
		}
	}

	return nil
}

func (f *File) String() string {
	return fmt.Sprintf(
		"File{Name:%s, Size: %d, Type: %v, Parent: %v, Files:%s}",
		f.Name, f.Size, f.Type, f.Parent, f.File,
	)
}

func PartOne(input []string) string {
	currentDir := NewFile("/", Directory, 0)

	for i := 1; i < len(input); i++ {

		fmt.Printf("currDir: %+v\n", currentDir)

		line := input[i]
		fmt.Println(line)

		// handle commands
		if strings.HasPrefix(line, "$ cd ") {
			dirName := strings.TrimPrefix(line, "$ cd ")

			newDir := currentDir.CD(dirName)
			fmt.Printf("a currDir: %+v\n", currentDir)

			newDir.Parent = currentDir
			fmt.Printf("b currDir: %+v\n", currentDir)
			currentDir = newDir

			fmt.Printf("c currDir: %+v\n", currentDir)
			continue
		}

		if strings.HasPrefix(line, "$ ls") {

			for i < len(input)-1 {
				i++
				line = input[i]

				if strings.HasPrefix(line, "$") {
					i--
					break
				}

				lineArr := strings.Fields(line)
				var newFile *File
				if lineArr[0] == "dir" {
					newFile = NewFile(lineArr[1], Directory, 0)
				} else {
					size, err := strconv.Atoi(lineArr[0])
					utils.CheckErr(err)

					newFile = NewFile(lineArr[1], RegularFile, size)
				}

				fmt.Println("1 AA")
				newFile.Parent = currentDir
				currentDir.File = append(currentDir.File, newFile)
			}
			fmt.Println("2 AA", currentDir)
		}
		fmt.Println("3 AA", currentDir)
	}

	return ""
}

func PartTwo(input []string) string {
	return ""
}
