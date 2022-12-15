package main

import (
	"fmt"
	"strings"
)

type ElfDirectory struct {
	name     string
	files    map[string]*ElfFile
	children map[string]*ElfDirectory
	parent   *ElfDirectory
}

func newElfDirectory(name string, parent *ElfDirectory) ElfDirectory {
	ed := ElfDirectory{
		name:   name,
		parent: parent,
		files:  make(map[string]*ElfFile),
	}

	if parent != nil {
		if parent.children == nil {
			parent.children = make(map[string]*ElfDirectory)
		}
		parent.children[name] = &ed
	}

	return ed
}

func (d *ElfDirectory) isParsable(input string) bool {
	return strings.HasPrefix(input, "dir")
}

func parseElfDirectory(input string, parent *ElfDirectory) *ElfDirectory {
	e := ElfDirectory{}

	if !e.isParsable(input) {
		return &e
	}

	slices := strings.Split(input, " ")

	name := slices[1]

	e = newElfDirectory(name, parent)

	return &e
}

func (d *ElfDirectory) toString() string {
	return fmt.Sprintf("ElfDirectory Name: %s", d.name)
}

func (d *ElfDirectory) addFile(f ElfFile) {
	if d.files == nil {
		d.files = make(map[string]*ElfFile)
	}
	d.files[f.name] = &f
}

func (d *ElfDirectory) size() int {
	size := 0
	for v := range d.files {
		size += d.files[v].size
	}

	for v := range d.children {
		child := d.children[v]
		size += child.size()
	}

	return size
}

func (d *ElfDirectory) findAllDirectories() []*ElfDirectory {
	var allDirs []*ElfDirectory

	for _, v := range d.children {
		allDirs = append(allDirs, v)
		allSubDirs := v.findAllDirectories()
		for i := 0; i < len(allSubDirs); i++ {
			allDirs = append(allDirs, allSubDirs[i])
		}
	}

	return allDirs
}
