package _024

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var maxY = 0
var maxX = 0
var corruptedBytes []Point

func orientedDijkstraDay18(Start Point, End Point, points Set[Point]) (int, map[int][][]Point) {
	var isValid = func(p Point) bool {
		return points.Contains(p)
	}

	var getPath = func(i *Item) []Point {
		var path []Point
		path = append(path, i.Point)
		for i.Predecessor != nil {
			i = i.Predecessor
			path = append(path, i.Point)
		}
		return path
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	distances := make(map[Point]map[Direction]int)
	for point := range points {
		distances[point] = map[Direction]int{
			UP:    math.MaxInt,
			RIGHT: math.MaxInt,
			DOWN:  math.MaxInt,
			LEFT:  math.MaxInt,
		}
	}

	initialDirection := LEFT
	heap.Push(pq, &Item{Point: Start, Direction: initialDirection, Cost: 0, Priority: 0})
	distances[Start][initialDirection] = 0

	cost := math.MaxInt
	paths := make(map[int][][]Point)
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		if current.Point == End {
			_, exists := paths[current.Cost]
			if !exists {
				paths[current.Cost] = make([][]Point, 0)
			}
			paths[current.Cost] = append(paths[current.Cost], getPath(current))
			if current.Cost < cost {
				cost = current.Cost
			}
		}

		for _, dir := range directions {
			neighbor := Point{current.Point.y + dir.Delta.y, current.Point.x + dir.Delta.x}
			if !isValid(neighbor) {
				continue
			}
			newCost := current.Cost + 1

			if newCost < distances[neighbor][dir.Direction] {
				distances[neighbor][dir.Direction] = newCost
				heap.Push(pq, &Item{
					Point:       neighbor,
					Direction:   dir.Direction,
					Predecessor: current,
					Cost:        newCost,
					Priority:    newCost,
				})
			}
		}
	}
	return cost, paths
}

func Day18() (int, string) {
	start := time.Now()
	parseInputDay18()
	part1, part2 := solutionDay18()
	elapsed := time.Since(start)
	fmt.Println("Solution day 18 time:", elapsed)
	return part1, part2
}

func parseInputDay18() {
	file, err := os.Open("input/2024/day18.txt")
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
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		corruptedByte := Point{y, x}
		corruptedBytes = append(corruptedBytes, corruptedByte)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
}

func solutionDay18() (int, string) {
	bytes := 1024
	byteSet := Set[Point]{}
	byteSet.UnionArray(corruptedBytes[:bytes])
	memoryLocations := Set[Point]{}
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			memoryLocations.Add(Point{i, j})
		}
	}
	memoryLocations = memoryLocations.difference(byteSet)
	memoryLocationsPart2 := memoryLocations.Copy()
	var lastByte = Point{0, 0}
	for i := 1024 + 1; i < 3500; i++ {
		memoryLocationsPart2.Remove(corruptedBytes[i])
		cost, _ := orientedDijkstraDay18(Point{0, 0}, Point{maxY, maxX}, memoryLocationsPart2)
		if cost > 1000000000 {
			lastByte = corruptedBytes[i]
			break
		}
	}
	cost, _ := orientedDijkstraDay18(Point{0, 0}, Point{maxY, maxX}, memoryLocations)
	lastByteString := fmt.Sprintf("%d,%d", lastByte.x, lastByte.y)
	return cost, lastByteString
}
