package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strings"
	"sync"
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
type State struct {
	loc Location
	dir Direction
}

var room [][]string
var room2 [][]string
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

func copyRoom(original [][]string) [][]string {
	newRoom := make([][]string, len(original))
	for i := range original {
		newRoom[i] = make([]string, len(original[i]))
		copy(newRoom[i], original[i])
	}
	return newRoom
}

func findGuard(room [][]string) (Location, bool) {
	for i := range room {
		for j := range room[i] {
			if room[i][j] == "^" {
				return Location{i, j}, true
			}
		}
	}
	return Location{-1, -1}, false
}

func simulateLoop(roomCopy [][]string, obstaclePos Location) bool {
	if obstaclePos.x >= 0 && obstaclePos.y >= 0 {
		roomCopy[obstaclePos.x][obstaclePos.y] = "#"
	}

	currentLoc, found := findGuard(roomCopy)
	if !found {
		return false
	}

	currentDir := Direction{-1, 0} // up
	visited := make(map[State]bool)

	for {
		nextX, nextY := currentLoc.x+currentDir.x, currentLoc.y+currentDir.y

		if nextX < 0 || nextX >= len(roomCopy) || nextY < 0 || nextY >= len(roomCopy[0]) {
			return false
		}

		currentState := State{currentLoc, currentDir}
		if visited[currentState] {
			return true
		}
		visited[currentState] = true

		if roomCopy[currentLoc.x][currentLoc.y] != "#" {
			roomCopy[currentLoc.x][currentLoc.y] = "X"
		}

		if roomCopy[nextX][nextY] == "#" {
			switch {
			case currentDir.x == -1 && currentDir.y == 0: // facing up
				currentDir = Direction{0, 1} // turn right
			case currentDir.x == 0 && currentDir.y == 1: // facing right
				currentDir = Direction{1, 0} // turn down
			case currentDir.x == 1 && currentDir.y == 0: // facing down
				currentDir = Direction{0, -1} // turn left
			case currentDir.x == 0 && currentDir.y == -1: // facing left
				currentDir = Direction{-1, 0} // turn up
			}
		} else {
			currentLoc = Location{nextX, nextY}
		}

		if len(visited) > len(roomCopy)*len(roomCopy[0])*4 {
			return false
		}
	}
}

/*
func countPossibleLoops(originalRoom [][]string) int {
	loopCount := 0

	for x := range originalRoom {
		for y := range originalRoom[x] {
			if originalRoom[x][y] == "." {
				roomCopy := copyRoom(originalRoom)
				if simulateLoop(roomCopy, Location{x, y}) {
					loopCount++
				}
			}
		}
	}

	return loopCount
}*/

func countPossibleLoops(originalRoom [][]string, startLine, endLine int) int {
	loopCount := 0

	for x := startLine; x < endLine; x++ {
		for y := range originalRoom[x] {
			if originalRoom[x][y] == "." {
				roomCopy := copyRoom(originalRoom)
				if simulateLoop(roomCopy, Location{x, y}) {
					loopCount++
				}
			}
		}
	}

	return loopCount
}

func worker(room [][]string, jobs <-chan [2]int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		startLine, endLine := job[0], job[1]
		results <- countPossibleLoops(room, startLine, endLine)
	}
}

func main() {
	path := "input.txt"
	room = readInput(path)
	room2 = readInput(path)
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

	visited := countXs()
	fmt.Println("Part 1:", visited, "in:", elapsed)

	// Part 2
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	jobs := make(chan [2]int, numCPU)
	results := make(chan int, numCPU)
	var wg sync.WaitGroup

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go worker(room2, jobs, results, &wg)
	}

	// Split the input data into ranges of lines
	chunkSize := (len(room2) + numCPU - 1) / numCPU
	start = time.Now()
	for i := 0; i < len(room2); i += chunkSize {
		end := i + chunkSize
		if end > len(room2) {
			end = len(room2)
		}
		jobs <- [2]int{i, end}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	loops := 0
	for result := range results {
		loops += result
	}
	elapsed = time.Since(start)

	fmt.Println("Part 2:", loops, "in:", elapsed)
}
