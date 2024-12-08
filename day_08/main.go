package main

import (
	"fmt"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	map_ := initializeMap(lines)
	antenna := extractAntennaPositions(lines)

	_, count1 := markAntinodes(map_, antenna, len(lines), len(lines[0]), calculateAntinode1)
	_, count2 := markAntinodes(map_, antenna, len(lines), len(lines[0]), calculateAntinode2)
	printMap(map_)
	fmt.Println("Part 1 : ", count1)
	fmt.Println("Part 2 : ", count2)
}

func initializeMap(lines []string) [][]string {
	map_ := make([][]string, len(lines))
	for i, line := range lines {
		map_[i] = make([]string, len(line))
		for j, val := range line {
			map_[i][j] = string(val)
		}
	}
	return map_
}

func extractAntennaPositions(lines []string) map[int32][]position {
	antenna := make(map[int32][]position)
	for i, line := range lines {
		for j, val := range line {
			if val != '.' {
				antenna[val] = append(antenna[val], position{i, j})
			}
		}
	}
	return antenna
}

func markAntinodes(map_ [][]string, antenna map[int32][]position, maxX, maxY int, findAntinode func(pos1, pos2 position, maxX, maxY int) []position) ([][]string, int) {
	count := 0
	mapCopy := copyMap(map_)
	for _, positions := range antenna {
		for i := range positions {
			for j := i + 1; j < len(positions); j++ {
				antinodes := findAntinode(positions[i], positions[j], maxX, maxY)
				for _, antinode := range antinodes {
					if mapCopy[antinode.x][antinode.y] != "#" {
						count++
						mapCopy[antinode.x][antinode.y] = "#"
					}
				}
			}
		}
	}
	return mapCopy, count
}

func printMap(map_ [][]string) {
	for _, line := range map_ {
		for _, val := range line {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func getDoubleDistance(pos1, pos2 position) position {
	return position{pos2.x + 2*(pos1.x-pos2.x), pos2.y + 2*(pos1.y-pos2.y)}
}

func getDistance(pos1, pos2 position, mult int) position {
	return position{pos2.x + mult*(pos1.x-pos2.x), pos2.y + mult*(pos1.y-pos2.y)}
}

func checkBounds(pos position, maxX, maxY int) bool {
	return pos.x >= 0 && pos.x < maxX && pos.y >= 0 && pos.y < maxY
}

func calculateAntinode1(pos1, pos2 position, maxX, maxY int) []position {
	var positions []position
	pos1Double := getDoubleDistance(pos1, pos2)
	pos2Double := getDoubleDistance(pos2, pos1)
	if checkBounds(pos1Double, maxX, maxY) {
		positions = append(positions, pos1Double)
	}
	if checkBounds(pos2Double, maxX, maxY) {
		positions = append(positions, pos2Double)
	}
	return positions
}

func calculateAntinode2(pos1, pos2 position, maxX, maxY int) []position {
	var positions []position
	positions = append(positions, calculateAntinodes(pos1, pos2, maxX, maxY)...)
	positions = append(positions, calculateAntinodes(pos2, pos1, maxX, maxY)...)
	return positions
}

func calculateAntinodes(start, end position, maxX, maxY int) []position {
	var antinodes []position
	val := 1
	for {
		antinode := getDistance(start, end, val)
		if !checkBounds(antinode, maxX, maxY) {
			break
		}
		antinodes = append(antinodes, antinode)
		val++
	}
	return antinodes
}

func copyMap(original [][]string) [][]string {
	copy_ := make([][]string, len(original))
	for i := range original {
		copy_[i] = make([]string, len(original[i]))
		copy(copy_[i], original[i])
	}
	return copy_
}
