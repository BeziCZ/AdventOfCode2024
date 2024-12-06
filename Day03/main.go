package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func findValidMul(line string, regex string) [][]string {
	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(line, -1)
	return matches
}

func part1(lines []string) int {
	start := time.Now()
	part1 := `mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)`
	res := 0
	matches := make([][]string, 0)

	for _, line := range lines {
		matches = append(matches, findValidMul(line, part1)...)
	}

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		res += x * y
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 took: ", elapsed)
	return res
}

func part2(lines []string) int {
	start := time.Now()
	part2 := `(?:mul\(([1-9]\d{0,2}),([1-9]\d{0,2})\)|do(?:n't)?\(\))`
	res := 0
	matches := make([][]string, 0)
	counting := true

	for _, line := range lines {
		matches = append(matches, findValidMul(line, part2)...)
	}

	for _, match := range matches {
		if match[0] == "don't()" {
			counting = false
		} else if !counting && match[0] == "do()" {
			counting = true
		}
		if counting {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			res += x * y
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 2 took: ", elapsed)
	return res
}

func main() {
	lines := readInput("input")

	fmt.Println("Result of Part 1 is:", part1(lines))
	fmt.Println("Result of Part 2 is:", part2(lines))

}
