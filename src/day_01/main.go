package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("src/day_01/input")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	var col1, col2 []int

	for _, line := range lines {
		fields := strings.Fields(line)

		val1, _ := strconv.Atoi(fields[0])
		val2, _ := strconv.Atoi(fields[1])
		col1 = append(col1, val1)
		col2 = append(col2, val2)
	}

	fmt.Println("Part 1 :", part1(col1, col2))
	fmt.Println("Part 2 :", part2(col1, col2))

}

func part1(col1 []int, col2 []int) int {
	sort.Ints(col1)
	sort.Ints(col2)
	var result int
	for i := range col1 {
		result += int(math.Abs(float64(col1[i] - col2[i])))
	}
	return result
}

func part2(col1 []int, col2 []int) int {
	occurrences := make(map[int]int)
	for _, val := range col2 {
		if slices.Contains(col1, val) {
			occurrences[val] += 1
		}
	}

	var result = 0
	for key, val := range occurrences {
		result += key * val
	}
	return result
}
