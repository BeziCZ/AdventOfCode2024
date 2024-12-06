package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		var row []int
		for _, s := range split {
			i, _ := strconv.Atoi(s)
			row = append(row, i)
		}
		result = append(result, row)
	}
	return result
}

func checkDescendency(arr []int) bool {
	for _, el := range arr {
		if el >= 0 {
			return false
		}
	}
	return true
}

func checkAscendancy(arr []int) bool {
	for _, el := range arr {
		if el <= 0 {
			return false
		}
	}
	return true
}

func isReportSafe(distances []int) bool {
	descOrAsc := checkDescendency(distances) || checkAscendancy(distances)
	allWithinBoundaries := true
	for _, dist := range distances {
		if !(math.Abs(float64(dist)) >= 1 && math.Abs(float64(dist)) <= 3) {
			allWithinBoundaries = false
			break
		}
	}
	return descOrAsc && allWithinBoundaries
}
func main() {
	matrix := readInput("/home/bezi/Projects/AdventOfCode/Day2/input.txt")

	numOfSafe1 := 0
	numOfSafe2 := 0
	// Part 1
	start := time.Now()
	for _, row := range matrix {
		distances := make([]int, 0)
		for i := 0; i < len(row)-1; i++ {
			distances = append(distances, int(row[i])-int(row[i+1]))
		}
		if isReportSafe(distances) {
			numOfSafe1++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 took: ", elapsed)
	// Part 2
	start = time.Now()
	for _, row := range matrix {
		distances := make([]int, 0)
		for i := 0; i < len(row)-1; i++ {
			distances = append(distances, int(row[i])-int(row[i+1]))
		}
		if isReportSafe(distances) {
			numOfSafe2++
		} else {
			for i := 0; i < len(row); i++ { // Go through all distances
				cpy_row := make([]int, len(row))
				mod_distances := make([]int, 0)
				copy(cpy_row, row)
				mod_row := append(cpy_row[:i], cpy_row[i+1:]...) // Remove element at index i
				for j := 0; j < len(mod_row)-1; j++ {
					mod_distances = append(mod_distances, int(mod_row[j])-int(mod_row[j+1]))
				}
				if isReportSafe(mod_distances) { // Check if modified distances are safe
					numOfSafe2++ // If they are, increment the number of safe reports
					break
				}
			}
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Part 2 took: ", elapsed)
	fmt.Println("Part 1 number of safe reports: ", numOfSafe1)
	fmt.Println("Part 2 number of safe reports: ", numOfSafe2)
}
