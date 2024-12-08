package main

import (
	"fmt"
	"os"
	"strings"
)

type direction struct {
	x int
	y int
}

var directions = []direction{{1, 0}, {0, 1}, {1, 1}, {-1, 1}}

func main() {
	count := 0
	data, err := os.ReadFile("src/day_04/input")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, dir := range directions {
		for i, line := range lines {
			for j := range line {
				if hasWord(lines, dir, i, j, "XMAS") {
					count++
				}
			}
		}

	}
	fmt.Println("Part 1 :", count)
	fmt.Println("Part 2 :", part2(lines))
}

func part2(lines []string) int {
	count := 0
	dirFirst := direction{1, 1}
	dirSecond := direction{1, -1}
	for i, line := range lines {
		for j := range line {
			if hasWord(lines, dirFirst, i, j, "MAS") {
				if j+2 < len(line) {
					if hasWord(lines, dirSecond, i, j+2, "MAS") {
						count++
					}
				}
			}

		}
	}
	return count
}

func hasWord(lines []string, dir direction, startX int, startY int, word string) bool {
	word_ := ""
	for k := range len(word) {
		if startX+(dir.x*(len(word)-1)) < 0 || startX+(dir.x*(len(word)-1)) >= len(lines) || startY+(dir.y*(len(word)-1)) < 0 || startY+(dir.y*(len(word)-1)) >= len(lines[0]) {
			break
		}
		word_ += string(lines[startX+(dir.x*k)][startY+(dir.y*k)])
	}

	if word_ == word || word_ == reverse(word) {
		return true
	}
	return false
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
