package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Day20() (int, int) {
	start := time.Now()
	parseInputDay20()
	part1, part2 := solutionDay20()
	elapsed := time.Since(start)
	fmt.Println("Solution day n time:", elapsed)
	return part1, part2
}

var start Point
var end Point
var track Set[Point]

func bfs(mazePoints Set[Point], start Point, end Point) []Point {
	var queue = Queue[Point]{}
	var visitedPoints = Set[Point]{}

	var parent = make(map[Point]Point)
	parent[start] = start

	queue.Enqueue(start)

	current, hasItem := queue.Dequeue()

	for hasItem {
		if current == end {
			break
		}
		for _, p := range current.getAdjacent() {
			if visitedPoints.Contains(p) || !mazePoints.Contains(p) {
				continue
			}
			queue.Enqueue(p)
			parent[p] = current
		}
		visitedPoints.Add(current)
		current, hasItem = queue.Dequeue()
	}

	var path []Point
	path = append(path, end)
	for current != parent[current] {
		path = append(path, parent[current])
		current = parent[current]
	}
	return path
}

func parseInputDay20() {
	file, err := os.Open("input/2024/day20.txt")
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
	counter := 0
	track = Set[Point]{}
	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		for index, char := range chars {
			switch char {
			case '.':
				track.Add(Point{counter, index})
				break
			case 'S':
				track.Add(Point{counter, index})
				start = Point{counter, index}
				break
			case 'E':
				track.Add(Point{counter, index})
				end = Point{counter, index}
				break
			}
			width = index
		}
		counter++
	}
	height = counter
}

func solutionDay20() (int, int) {
	path := bfs(track, start, end)
	part1 := findCheats(path, 2, 100)
	part2 := findCheats(path, 20, 100)
	return part1, part2
}

func findCheats(points []Point, pathLength int, saveTime int) int {
	sum := 0
	saveTime++
	for i, point := range points {
		if saveTime+i >= len(points) {
			break
		}
		for skipIndex, skipPoint := range points[saveTime+i:] {
			skipLength := point.manhattan(skipPoint)
			if skipLength <= pathLength && skipIndex-skipLength+1 >= 0 {
				sum++
			}
		}
	}
	return sum
}

func printPath(path []Point) {
	pathSet := Set[Point]{}
	pathSet.UnionArray(path)
	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			point := Point{i, j}
			if point == start {
				print("S")
			} else if point == end {
				print("E")
			} else if pathSet.Contains(point) {
				print("\033[33mO\033[0m")
			} else if track.Contains(point) {
				print("\033[36m.\033[0m")
			} else {
				print("#")
			}
		}
		println("")
	}
}
