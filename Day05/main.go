package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
func buildGraph(pairs [][]int) (map[int][]int, map[int]int, map[int]bool) {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	nodes := make(map[int]bool)

	for _, pair := range pairs {
		first, second := pair[0], pair[1]
		if _, exists := graph[first]; !exists {
			graph[first] = []int{}
		}
		graph[first] = append(graph[first], second)
		nodes[first] = true
		nodes[second] = true
		inDegree[second]++
	}
	fmt.Println("Graph:", graph)
	fmt.Println("InDegree:", inDegree)
	fmt.Println("Nodes:", nodes)
	return graph, inDegree, nodes
}

func topologicalSort(pairs [][]int) []int {
	graph, inDegree, nodes := buildGraph(pairs)

	// Find nodes with no incoming edges
	var queue []int
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)

		}
	}

	var result []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}
*/

type Rule struct {
	a, b int
}

func isUpdateValid(line []int, rules []Rule) bool {
	for _, rule := range rules {
		aPos, bPos := -1, -1

		// Find positions of both a and b in the line
		for i, page := range line {
			if page == rule.a {
				aPos = i
			}
			if page == rule.b {
				bPos = i
			}
		}

		// If both a and b are present and a comes after b, it's invalid
		if aPos != -1 && bPos != -1 && aPos > bPos {
			return false
		}
	}
	return true
}

func topologicalSort(line []int, rules []Rule) []int {
	sorted := false
	for !sorted {
		sorted = true
		for _, rule := range rules {
			// If the current rule is violated, swap the positions of the pages
			for i := 0; i < len(line)-1; i++ {
				for j := i + 1; j < len(line); j++ {
					if line[i] == rule.b && line[j] == rule.a {
						// Swap the pages
						line[i], line[j] = line[j], line[i]
						sorted = false
					}
				}
			}
		}
	}
	return line
}

func readInput(path string) ([][]int, []Rule) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var updates [][]int
	var rules []Rule

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			split := strings.Split(line, "|")
			first, _ := strconv.Atoi(split[0])
			second, _ := strconv.Atoi(split[1])
			rules = append(rules, Rule{first, second})
		} else {
			split := strings.Split(line, ",")
			var update []int
			for _, num := range split {
				n, _ := strconv.Atoi(num)
				update = append(update, n)
			}
			updates = append(updates, update)
		}
	}
	return updates, rules
}

func main() {
	updates, rules := readInput("input.txt")
	result1 := 0

	// Part 1
	start := time.Now()
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			result1 += update[len(update)/2]
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 result:", result1, "in", elapsed)

	// Part 2
	result2 := 0
	start = time.Now()
	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			sorted := topologicalSort(update, rules)
			result2 += sorted[len(sorted)/2]
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Part 2 result:", result2, "in", elapsed)
}
