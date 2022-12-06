package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fmt.Println(findMarker(line, 4))
	fmt.Println(findMarker(line, 14))
}

func findMarker(line string, length int) int {
	seen := map[byte]struct{}{}
	var left, right = 0, 0

	// move `right` forward
	for ; ; right++ {

		// if we find a duplicate character
		if _, ok := seen[line[right]]; ok {

			// we move `left` forward, up to the next
			// of the first occurrence of the duplicate
			for ; ; left++ {

				if line[left] == line[right] {
					left++
					break
				}
				delete(seen, line[left])
			}

			// no need to add on `seen` (we didn't delete it)
			// no need to check the length cause moving left
			// shorten the string
			continue
		}

		seen[line[right]] = struct{}{}

		// if [left,right) is long enough, that's our goal
		if right-left+1 == length {
			return right + 1
		}
	}
}

