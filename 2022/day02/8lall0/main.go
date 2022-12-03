package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part 1")
	n, err := part1(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Score: %d\n", n)

	fmt.Println("Part 2")
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n, err = part2(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Score: %d", n)
}
