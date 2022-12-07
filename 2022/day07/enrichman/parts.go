package main

import (
	"strings"
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

func PartOne(input []string) string {
	currentDir := NewFile("/", Directory, 0)

	for i := 0; i < len(input); i++ {
		// handle commands
		if strings.HasPrefix(line, "$ cd") {
			dirName := strings.TrimPrefix(s, "$ cd")
			if dirName == ".." {
				currentDir = currentDir.Parent
				continue
			}

			newDir := NewFile(dirName, Directory, 0)
			newDir.Parent = currentDir
			currentDir = newDir
			continue
		}
	}

	return ""
}

func PartTwo(input []string) string {
	return ""
}

type Cmd interface {
	Execute(currentDir *File)
}

type ChangeDirCmd struct{}

func (cd *ChangeDirCmd) Execute(currentDir *File) {

}

type ListCmd struct{}

func (ls *ListCmd) Execute(currentDir *File) {
	for _, f := range currentDir.File {

	}
}

func parseCommand(s string) Cmd {
	cmd := strings.TrimPrefix(s, "$ ")
	switch cmd {
	case strings.HasPrefix("cd"):
		c = &Cmd{}
	case strings.HasPrefix("ls"):
		c = &Cmd{}
	}
}
