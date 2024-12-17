package _024

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type MazePoint struct {
	point     Point
	direction Direction
}

func Day16() (int, int) {
	start := time.Now()
	parseInputDay16()
	part1, part2 := solutionDay16()
	elapsed := time.Since(start)
	fmt.Println("Solution day 16 time:", elapsed)
	return part1, part2
}

var pathPoints = Set[Point]{}
var Start MazePoint
var End Point

func parseInputDay16() {
	file, err := os.Open("input/2024/day16.txt")
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
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		row := []rune(line)
		for i, v := range row {
			if v == '.' || v == 'S' || v == 'E' {
				pathPoints.Add(Point{counter, i})
			}
			if v == 'S' {
				Start = MazePoint{Point{counter, i}, RIGHT}
			}
			if v == 'E' {
				End = Point{counter, i}
			}
			width = i
		}
		counter++
	}
	height = counter
}

func orientedDijkstra(Start MazePoint, End Point, points Set[Point]) (int, map[int][][]Point) {
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
	for point := range pathPoints {
		distances[point] = map[Direction]int{
			UP:    math.MaxInt,
			RIGHT: math.MaxInt,
			DOWN:  math.MaxInt,
			LEFT:  math.MaxInt,
		}
	}

	initialDirection := LEFT
	heap.Push(pq, &Item{Point: Start.point, Direction: initialDirection, Cost: 0, Priority: 0})
	distances[Start.point][initialDirection] = 0

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
			moveCost := 1
			turnCost := 0
			if current.Direction != dir.Direction {
				turnCost = 1000
			}
			newCost := current.Cost + moveCost + turnCost

			if newCost <= distances[neighbor][dir.Direction] {
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

func solutionDay16() (int, int) {
	part1, scorePaths := orientedDijkstra(Start, End, pathPoints)

	points := Set[Point]{}
	for score, paths := range scorePaths {
		if score != part1 {
			continue
		}
		for _, path := range paths {
			for _, point := range path {
				points.Add(point)
			}
		}
	}
	return part1, len(points.Items())
}

func printPaths(routePoints Set[Point]) {
	for i := 0; i < height; i++ {
		for j := 0; j <= width; j++ {
			if routePoints.Contains(Point{i, j}) {
				print("O")
			} else if pathPoints.Contains(Point{i, j}) {
				print(".")
			} else {
				print("#")
			}
		}
		println("")
	}
}
