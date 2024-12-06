package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		result = append(result, split)
	}
	return result
}

func countWord(input [][]string, word string) int {
	height := len(input)
	width := len(input[0])
	count := 0
	wordLength := len(word)

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			for _, dir := range directions {
				if input[row][col] != string(word[0]) {
					continue
				}

				match := true
				for i := 0; i < wordLength; i++ {
					currRow := row + (i * dir[0])
					currCol := col + (i * dir[1])
					if currRow < 0 || currRow >= height || currCol < 0 || currCol >= width || input[currRow][currCol] != string(word[i]) {
						match = false
						break
					}
				}
				if match {
					count++
				}
			}
		}
	}
	return count
}

func countMAS(input [][]string) int {
	height := len(input)
	width := len(input[0])
	count := 0

	// We need at least a 3x3 space to check the pattern
	for row := 0; row <= height-3; row++ {
		for col := 0; col <= width-3; col++ {
			if input[row+1][col+1] != "A" {
				continue
			}

			if input[row][col] == "M" &&
				input[row][col+2] == "S" &&
				input[row+2][col] == "M" &&
				input[row+2][col+2] == "S" {
				count++
			}

			if input[row][col] == "S" &&
				input[row][col+2] == "M" &&
				input[row+2][col] == "S" &&
				input[row+2][col+2] == "M" {
				count++
			}

			if input[row][col] == "S" &&
				input[row][col+2] == "S" &&
				input[row+2][col] == "M" &&
				input[row+2][col+2] == "M" {
				count++
			}

			if input[row][col] == "M" &&
				input[row][col+2] == "M" &&
				input[row+2][col] == "S" &&
				input[row+2][col+2] == "S" {
				count++
			}
		}
	}
	return count
}

func main() {
	input := readInput("input.txt")

	start := time.Now()
	count := countWord(input, "XMAS")
	elapse := time.Since(start)
	fmt.Println("Part 1 total:", count, "found in:", elapse)

	start = time.Now()
	count = countMAS(input)
	elapse = time.Since(start)
	fmt.Println("Part 2 total:", count, "found in:", elapse)
}
