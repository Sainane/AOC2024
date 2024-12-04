package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	corruptedMemory, err := os.ReadFile("input")

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(string(corruptedMemory), -1)

	totalSum := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		totalSum += x * y
	}

	fmt.Println("Sum of all valid mul instructions:", totalSum)
	fmt.Println("Sum of all valid mul instructions part 2:", part2(string(corruptedMemory)))
}

func part2(corruptedMemory string) int {
	totalSum := 0
	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	stateRegex := regexp.MustCompile(`do\(\)|don't\(\)`)
	enabled := true
	pos := 0
	for pos < len(corruptedMemory) {
		mulMatch := mulRegex.FindStringSubmatchIndex(corruptedMemory[pos:])
		stateMatch := stateRegex.FindStringIndex(corruptedMemory[pos:])

		if mulMatch == nil && stateMatch == nil {
			break
		}

		if mulMatch != nil && (stateMatch == nil || mulMatch[0] < stateMatch[0]) {

			x, _ := strconv.Atoi(corruptedMemory[pos+mulMatch[2] : pos+mulMatch[3]])
			y, _ := strconv.Atoi(corruptedMemory[pos+mulMatch[4] : pos+mulMatch[5]])

			if enabled {
				totalSum += x * y
			}

			pos += mulMatch[1]
		} else if stateMatch != nil {
			stateInstruction := corruptedMemory[pos+stateMatch[0] : pos+stateMatch[1]]
			if stateInstruction == "do()" {
				enabled = true
			} else if stateInstruction == "don't()" {
				enabled = false
			}

			pos += stateMatch[1]
		}
	}
	return totalSum

}
