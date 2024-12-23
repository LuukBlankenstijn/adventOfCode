package _024

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
	"unicode"
)

type cacheKey struct {
	start Point
	end   Point
	layer int
}

var day21Cache = make(map[cacheKey]int64)

var inputCodes []string

var dirToString = map[Direction]string{
	UP:    "^",
	DOWN:  "v",
	LEFT:  "<",
	RIGHT: ">",
}

var numKeyPad = map[string]Point{
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"0": {3, 1},
	"A": {3, 2},
}
var dirPad = map[string]Point{
	"^": {0, 1},
	"A": {0, 2},
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
}

func Day21() (int64, int64) {
	start := time.Now()
	parseInputDay21()
	part1, part2 := solutionDay21()
	elapsed := time.Since(start)
	fmt.Println("Solution day n time:", elapsed)
	return part1, part2
}

func parseInputDay21() {
	file, err := os.Open("input/2024/day21.txt")
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
	for scanner.Scan() {
		line := scanner.Text()
		inputCodes = append(inputCodes, line)
	}
}

func solutionDay21() (int64, int64) {
	var sum int64
	var sumPart2 int64
	sum = 0
	sumPart2 = 0
	for _, code := range inputCodes {
		sum += getCodeLength(code, 2) * int64(codeToNumber(code))
	}
	day21Cache = map[cacheKey]int64{}
	for _, code := range inputCodes {
		sumPart2 += getCodeLength(code, 25) * int64(codeToNumber(code))
	}
	return sum, sumPart2
}

func codeToNumber(code string) int {
	output := ""
	for _, char := range code {
		if unicode.IsDigit(char) {
			output += string(char)
		}
	}
	number, _ := strconv.Atoi(output)
	return number
}

func getCodeLength(code string, endLayer int) int64 {
	numKeyPadArray := []Point{}
	for _, v := range numKeyPad {
		numKeyPadArray = append(numKeyPadArray, v)
	}
	code = "A" + code
	var sum int64
	sum = 0
	for i := 0; i < len(code)-1; i++ {
		start = numKeyPad[string(code[i])]
		end = numKeyPad[string(code[i+1])]
		sum += adfasdf(start, end, numKeyPadArray, 0, endLayer)
	}
	return sum
}

func adfasdf(start, end Point, points []Point, layer int, endLayer int) int64 {
	value, exists := day21Cache[cacheKey{start, end, layer}]
	if exists {
		return value
	}
	paths := bfsDay21(start, end, points)
	if layer == endLayer {
		return int64(len(paths[0]))
	}
	dirPadArray := []Point{}
	for _, v := range dirPad {
		dirPadArray = append(dirPadArray, v)
	}
	var minimum int64
	minimum = math.MaxInt64
	for _, path := range paths {
		path = "A" + path
		var sum int64
		sum = 0
		for i := 0; i < len(path)-1; i++ {
			sum += adfasdf(dirPad[string(path[i])], dirPad[string(path[i+1])], dirPadArray, layer+1, endLayer)
		}
		if sum < minimum {
			minimum = sum
		}
	}
	day21Cache[cacheKey{start, end, layer}] = minimum
	return minimum
}

func bfsDay21(start, end Point, points []Point) []string {
	visitedPoints := Set[Point]{}

	// Define bfsPoint struct to store the current point and its path point
	type bfsPoint struct {
		point Point
		path  string
	}

	// Initialize queue
	queue := []bfsPoint{{point: start, path: ""}}
	shortestPaths := []string{}
	shortestDistance := -1

	// BFS loop
	for len(queue) > 0 {
		// Dequeue
		current := queue[0]
		queue = queue[1:]

		// If we reached the end point
		if current.point == end {
			// Add to paths if it matches the shortest distance
			if shortestDistance == -1 || len(current.path) == shortestDistance {
				shortestPaths = append(shortestPaths, current.path+"A")
				shortestDistance = len(current.path)
			} else if len(current.path) < shortestDistance {
				// Found a shorter path, clear path paths
				shortestPaths = []string{current.path + "A"}
				shortestDistance = len(current.path)
			}
			continue
		}

		// Explore neighbors (up, down, left, right)
		for dir, neighbor := range current.point.getAdjacent() {
			// Check if the neighbor is valid and not visitedPoints
			if isPointValid(neighbor, points) && !visitedPoints.Contains(neighbor) {
				queue = append(queue, bfsPoint{
					point: neighbor,
					path:  current.path + dirToString[dir],
				})
			}
		}
		visitedPoints.Add(current.point)
	}

	return shortestPaths
}

// Helper function to check if a point is valid
func isPointValid(p Point, points []Point) bool {
	for _, validPoint := range points {
		if p == validPoint {
			return true
		}
	}
	return false
}
