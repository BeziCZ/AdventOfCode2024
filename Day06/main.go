package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Direction struct {
	x int
	y int
}
type Location struct {
	x int
	y int
}

var room [][]string
var currDir Direction = Direction{-1, 0}
var guardLoc Location

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error reading file")
	}
	defer file.Close()

	var input [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		input = append(input, split)
	}
	return input
}

func countXs() int {
	count := 0
	for _, row := range room {
		for _, cell := range row {
			if cell == "X" {
				count++
			}
		}
	}
	return count
}

func changeDirection() {
	if currDir.x == -1 && currDir.y == 0 {
		// Up -> Right
		currDir = Direction{0, 1}
	} else if currDir.x == 0 && currDir.y == 1 {
		// Right -> Down
		currDir = Direction{1, 0}
	} else if currDir.x == 1 && currDir.y == 0 {
		// Down -> Left
		currDir = Direction{0, -1}
	} else if currDir.x == 0 && currDir.y == -1 {
		// Left -> Up
		currDir = Direction{-1, 0}
	}
}

func makeMove() int {
	// Check if the next move is off the map
	if guardLoc.x+currDir.x < 0 || guardLoc.x+currDir.x >= len(room) || guardLoc.y+currDir.y < 0 || guardLoc.y+currDir.y >= len(room[0]) {
		room[guardLoc.x][guardLoc.y] = "X"
		return -1 // Guard is off the map
	}
	if room[guardLoc.x+currDir.x][guardLoc.y+currDir.y] == "#" {
		room[guardLoc.x][guardLoc.y] = "X"
		changeDirection()
		room[guardLoc.x+currDir.x][guardLoc.y+currDir.y] = "^"
		guardLoc = Location{guardLoc.x + currDir.x, guardLoc.y + currDir.y}
	} else {
		room[guardLoc.x][guardLoc.y] = "X"
		room[guardLoc.x+currDir.x][guardLoc.y+currDir.y] = "^"
		guardLoc = Location{guardLoc.x + currDir.x, guardLoc.y + currDir.y}
	}
	return 0
}

func main() {
	path := "input.txt"
	room = readInput(path)

	// Part 1
	// Find the guard's init location
	for x, row := range room {
		if slices.Contains(row, "^") {
			guardLoc = Location{x, slices.Index(row, "^")}
			break
		}
	}

	cond := 0
	start := time.Now()
	for ok := true; ok; ok = (cond == 0) {
		cond = makeMove()
	}
	elapsed := time.Since(start)

	for _, row := range room {
		fmt.Println(row)
	}
	visited := countXs()
	fmt.Println("Part 1: ", visited, "in:", elapsed)
}
