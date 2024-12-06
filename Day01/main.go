package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parseFile(path string) (array1 []int, array2 []int) {
	fmt.Println("Parsing file")
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Initialize arrays to hold the numbers
	var leftNumbers []int
	var rightNumbers []int

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Split the line into two parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("invalid line format: %s", line)
		}

		// Convert the parts to integers
		leftNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("failed to convert left number: %s", err)
		}
		rightNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to convert right number: %s", err)
		}

		// Append the numbers to the arrays
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}
	return leftNumbers, rightNumbers
}

func freq_map(array []int) map[int]int {
	freq := make(map[int]int)
	for _, value := range array {
		freq[value]++
	}
	return freq
}

func main() {
	fmt.Println("Starting to solve Day 1 of Advent of Code 2024")

	leftNumbers, rightNumbers := parseFile("/home/bezi/Projects/AdventOfCode/Day1/input.txt")

	start := time.Now()
	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	sum := 0

	for i := 0; i < len(leftNumbers); i++ {
		if leftNumbers[i] < rightNumbers[i] {
			sum += rightNumbers[i] - leftNumbers[i]
		} else {
			sum += leftNumbers[i] - rightNumbers[i]
		}
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time to do part 1: ", elapsed)
	start = time.Now()
	similarity := 0

	rightFreq := freq_map(rightNumbers)
	for i := 0; i < len(leftNumbers); i++ {
		similarity += rightFreq[leftNumbers[i]] * leftNumbers[i]
	}
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Time to do part 2: ", elapsed)
	fmt.Println("The part 1 result is: ", sum)
	fmt.Println("The part 2 result is: ", similarity)
}
