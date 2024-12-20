package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Equation struct {
	result  int
	numbers []int
}

func main() {
	data, err := os.ReadFile("src/day_07/input")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	var equations []Equation
	for _, line := range lines {
		val := strings.Split(line, ": ")
		res := strings.Split(val[1], " ")
		eq := Equation{}
		eq.result, _ = strconv.Atoi(val[0])
		for _, r := range res {
			numVal, _ := strconv.Atoi(r)
			eq.numbers = append(eq.numbers, numVal)
		}

		equations = append(equations, eq)
	}

	fmt.Println("Part 1 : ", solutionsForOp(equations, []func(a int, b int) int{plus, times}))
	fmt.Println("Part 2 : ", solutionsForOp(equations, []func(a int, b int) int{plus, times, or}))
}

func solutionsForOp(equations []Equation, ops_ []func(a int, b int) int) int {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	sol := 0

	for _, e := range equations {
		wg.Add(1)
		go func(e Equation) {
			defer wg.Done()
			ops := allOpPoss(len(e.numbers)-1, ops_)
			for _, op := range ops {
				res := e.numbers[0]
				for i, o := range op {
					res = o(res, e.numbers[i+1])
				}
				if res == e.result {
					mu.Lock()
					sol += e.result
					mu.Unlock()
					break
				}
			}
		}(e)
	}

	wg.Wait()
	return sol
}

func plus(a int, b int) int {
	return a + b
}

func times(a int, b int) int {
	return a * b
}

func or(a int, b int) int {
	unit := int(math.Log10(float64(b))) + 1
	return a*int(math.Pow(10.0, float64(unit))) + b
}

func allOpPoss(nbOp int, posOp []func(a int, b int) int) [][]func(a int, b int) int {
	var combinations [][]func(a int, b int) int
	generateOperatorCombinations(make([]func(a int, b int) int, nbOp), 0, &combinations, posOp)
	return combinations
}

func generateOperatorCombinations(operators []func(a int, b int) int, index int, combinations *[][]func(a int, b int) int, posOp []func(a int, b int) int) {
	if index == len(operators) {
		opCopy := make([]func(a int, b int) int, len(operators))
		copy(opCopy, operators)
		*combinations = append(*combinations, opCopy)
		return
	}

	for _, op := range posOp {
		operators[index] = op
		generateOperatorCombinations(operators, index+1, combinations, posOp)
	}
}
