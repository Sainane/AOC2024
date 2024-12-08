package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func parseInput(filename string) ([]Rule, [][]int, []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var rule []string
	var rules []Rule
	var updates [][]int
	scanner := bufio.NewScanner(file)

	// Parse rules
	parsingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			after, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			rules = append(rules, Rule{before, after})
			rule = append(rule, line)
		} else {
			var update []int
			for _, numStr := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(strings.TrimSpace(numStr))
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}
	return rules, updates, rule
}

func isValidOrder(update []int, rules []Rule) bool {
	// Create a map of indices for quick lookup
	indices := make(map[int]int)
	for i, num := range update {
		indices[num] = i
	}

	// Check each rule
	for _, rule := range rules {
		beforeIdx, beforeExists := indices[rule.before]
		afterIdx, afterExists := indices[rule.after]

		// Only check if both numbers exist in the update
		if beforeExists && afterExists {
			if beforeIdx > afterIdx {
				return false
			}
		}
	}
	return true
}

func main() {
	rules, updates, _ := parseInput("src/day_05/input")

	sum := 0
	for _, update := range updates {
		valid := isValidOrder(update, rules)
		if valid {
			middleIdx := len(update) / 2
			middleNum := update[middleIdx]
			sum += middleNum
		}
	}

	fmt.Println("Part 1 :", sum)
	part2()
}

type ordering struct {
	order int
	index int
}

func (o ordering) Less(other ordering) bool {
	if o.order < other.order {
		return true
	}
	if o.order == other.order {
		return o.index < other.index
	}

	return false
}

// Parse the rules into a map: page -> set of pages that must come after this page
func parseRules(rules []string) map[int]map[int]bool {
	ordering := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			continue
		}
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		if _, exists := ordering[before]; !exists {
			ordering[before] = make(map[int]bool)
		}
		ordering[before][after] = true
	}
	return ordering
}

func reorderUpdate(update []int, ordering map[int]map[int]bool) []int {
	indegree := make(map[int]int)
	graph := make(map[int][]int)

	// Create graph and compute degrees
	for _, page := range update {
		indegree[page] = 0
		graph[page] = []int{}
	}

	for before, afterSet := range ordering {
		if _, exists := indegree[before]; !exists {
			continue
		}
		for after := range afterSet {
			if _, exists := indegree[after]; !exists {
				continue
			}
			graph[before] = append(graph[before], after)
			indegree[after]++
		}
	}

	var queue []int
	for page, deg := range indegree {
		if deg == 0 {
			queue = append(queue, page)
		}
	}

	var sorted []int
	for len(queue) > 0 {
		page := queue[0]
		queue = queue[1:]
		sorted = append(sorted, page)

		for _, neighbor := range graph[page] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

// Compute the middle page index
func middlePage(update []int) int {
	mid := len(update) / 2
	return update[mid]
}

func part2() {

	_, updates, rules := parseInput("src/day_05/input")
	ordering := parseRules(rules)
	incorrectMiddleSum := 0

	for _, update := range updates {
		if !isCorrectOrder(update, ordering) {
			fixedUpdate := reorderUpdate(update, ordering)
			incorrectMiddleSum += middlePage(fixedUpdate)
		}
	}

	fmt.Println("Part 2 :", incorrectMiddleSum)
}

// Check if an update is in the correct order
func isCorrectOrder(update []int, ordering map[int]map[int]bool) bool {
	updateSet := make(map[int]int)
	for i, page := range update {
		updateSet[page] = i
	}
	for before, afterSet := range ordering {
		posBefore, existsBefore := updateSet[before]
		if !existsBefore {
			continue
		}
		for after := range afterSet {
			posAfter, existsAfter := updateSet[after]
			if existsAfter && posBefore >= posAfter {
				return false
			}
		}
	}
	return true
}
