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
	defer file.Close()
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
	rules, updates, _ := parseInput("input")

	// Debug output
	fmt.Println("Rules:")
	for _, r := range rules {
		fmt.Printf("%d must come before %d\n", r.before, r.after)
	}

	sum := 0
	fmt.Println("\nChecking updates:")
	for i, update := range updates {
		valid := isValidOrder(update, rules)
		if valid {
			middleIdx := len(update) / 2
			middleNum := update[middleIdx]
			sum += middleNum
			fmt.Printf("Update %d: %v is valid. Middle number: %d\n", i+1, update, middleNum)
		} else {
			fmt.Printf("Update %d: %v is invalid\n", i+1, update)
		}
	}

	fmt.Printf("\nFinal sum of middle numbers: %d\n", sum)
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

type toOrder []ordering

func (o toOrder) Len() int {
	return len(o)
}

func (o toOrder) Less(i, j int) bool {
	return o[i].Less(o[j])
}

func (o toOrder) Swap(i, j int) {
	temp := o[i]
	o[i] = o[j]
	o[j] = temp
}

// Utility function to convert string to int
func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Parse the rules into a map: page -> set of pages that must come after this page
func parseRules(rules []string) map[int]map[int]bool {
	ordering := make(map[int]map[int]bool)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		if len(parts) != 2 {
			continue
		}
		before := atoi(parts[0])
		after := atoi(parts[1])

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

	// Create graph and compute indegrees
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

	queue := []int{}
	for page, deg := range indegree {
		if deg == 0 {
			queue = append(queue, page)
		}
	}

	sorted := []int{}
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

	_, updates, rules := parseInput("input")
	ordering := parseRules(rules)
	incorrectMiddleSum := 0

	for _, update := range updates {
		if !isCorrectOrder(update, ordering) {
			fixedUpdate := reorderUpdate(update, ordering)
			incorrectMiddleSum += middlePage(fixedUpdate)
		}
	}

	fmt.Println("The sum of middle pages of correctly ordered (fixed) updates is:", incorrectMiddleSum)
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
