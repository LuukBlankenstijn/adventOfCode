package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var patterns []string
var towels []string
var towelMapping = make(map[string]int)

func Day19() (int, int) {
	start := time.Now()
	parseInputDay19()
	part1, part2 := solutionDay19()
	elapsed := time.Since(start)
	fmt.Println("Solution day 19 time:", elapsed)
	return part1, part2
}

func parseInputDay19() {
	file, err := os.Open("input/2024/day19.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	patterns = strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		towels = append(towels, line)
	}
}

func isTowelPossible(towel string) int {
	if value, exists := towelMapping[towel]; exists {
		return value
	}
	ways := 0
	for _, pattern := range patterns {
		if !strings.HasPrefix(towel, pattern) {
			continue
		}
		if len(towel) == len(pattern) {
			ways++
			continue
		}
		possibleWays := isTowelPossible(strings.Replace(towel, pattern, "", 1))
		ways += possibleWays
	}
	towelMapping[towel] = ways
	return ways
}

func solutionDay19() (int, int) {
	countPart1 := 0
	countPart2 := 0
	for _, towel := range towels {
		ways := isTowelPossible(towel)
		countPart2 += ways
		if ways > 0 {
			countPart1++
		}
	}
	return countPart1, countPart2
}
