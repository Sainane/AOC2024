package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	var values [][]int
	for _, line := range lines {
		fields := strings.Fields(line)
		var valLine []int
		for _, val := range fields {
			num, _ := strconv.Atoi(val)
			valLine = append(valLine, num)
		}
		values = append(values, valLine)

	}

	fmt.Println(part1(values))
	fmt.Println(part2(values))
}

func checkValidity(num []int) bool {
	increasing := num[0] < num[1]
	for i := range num {
		if i != len(num)-1 {
			smaller := num[i] < num[i+1]
			if smaller && !increasing || !smaller && increasing {

				return false
			}

			diff := math.Abs(float64(num[i] - num[i+1]))
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}

func checkValidity2(num []int) bool {
	validity := false
	for i := range num {
		var newTab []int
		for j, val := range num {
			if j != i {
				newTab = append(newTab, val)
			}
		}
		validity = checkValidity(newTab)

		if validity {
			return true
		}
	}
	return false
}
func part1(values [][]int) int {
	nbValid := 0
	for _, num := range values {
		if checkValidity(num) {
			nbValid++
		}
	}
	return nbValid

}

func part2(values [][]int) int {
	nbValid := 0
	for _, num := range values {
		if checkValidity2(num) {
			nbValid++
		}
	}
	return nbValid
}
