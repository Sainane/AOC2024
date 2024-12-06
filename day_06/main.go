package main

import (
	"bufio"
	"fmt"
	"os"
)

type MoveResult int

const (
	OUT MoveResult = iota
	OK
	STUCK
	LOOP
)

type Direction struct {
	x, y int
}

func (d Direction) equals(other Direction) bool {
	return d.x == other.x && d.y == other.y
}

func (d Direction) isOpposite(other Direction) bool {
	return d.x+other.x+d.y+other.y == 0
}

type Position struct {
	x, y int
}

func (p Position) move(d Direction) Position {
	return Position{p.x + d.x, p.y + d.y}
}

type Result struct {
	nbPos  int
	isLoop bool
}

var directions = []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func checkMove(grid [][]byte, pos Position, d Direction, nbIter int) MoveResult {
	newPos := pos.move(d)
	if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(grid) || newPos.y >= len(grid[0]) {
		return OUT
	}
	if grid[newPos.x][newPos.y] == '#' {
		return STUCK
	}
	if nbIter > len(grid)*len(grid[0]) {
		return LOOP
	}
	return OK
}

func moveInMap(grid [][]byte, pos Position, d Direction) Result {
	nbPos := 1
	dirIndex := 0
	nbIter := 0
	d = directions[dirIndex]

	for {
		result := checkMove(grid, pos, d, nbIter)
		switch result {
		case OUT:
			return Result{nbPos, false}
		case OK:
			pos = pos.move(d)
			nbIter++
			if grid[pos.x][pos.y] == '.' {
				grid[pos.x][pos.y] = 'X'
				nbPos++
			}
		case STUCK:
			dirIndex = (dirIndex + 1) % 4
			d = directions[dirIndex]
		case LOOP:
			return Result{nbPos, true}
		}
	}
}

func copyGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var grid [][]byte
	var startPos Position

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	// Find starting position
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' {
				startPos = Position{i, j}
				grid[i][j] = 'X'
			}
		}
	}

	// Part 1
	result := moveInMap(copyGrid(grid), startPos, directions[0])
	fmt.Printf("Part 1: %d\n", result.nbPos)

	// Part 2
	nbLoop := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				newGrid := copyGrid(grid)
				newGrid[i][j] = '#'
				result := moveInMap(newGrid, startPos, directions[0])
				if result.isLoop {
					nbLoop++
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", nbLoop)
}
