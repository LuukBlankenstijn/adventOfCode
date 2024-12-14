package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var garden [][]rune

var visited Set[Point]
var plots []Set[Point]

func Day12() (int, int) {
	//os.Exit(-1)
	start := time.Now()
	parseInputDay12()
	part1, part2 := solutionDay12()
	elapsed := time.Since(start)
	fmt.Println("Solution day 12 time:", elapsed)
	return part1, part2
}

func parseInputDay12() {
	file, err := os.Open("input/2024/day12.txt")
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
		row := []rune(line)
		garden = append(garden, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solutionDay12() (int, int) {
	visited = Set[Point]{}
	for yI, row := range garden {
		for xI, _ := range row {
			point := Point{yI, xI}
			if visited.Contains(point) {
				continue
			}
			patch := Set[Point]{}
			patch.Add(point)
			getPlot(point, patch)
			if len(patch.Items()) > 0 {
				visited.Union(patch)
				plots = append(plots, patch)
			}
		}
	}

	totalCost := 0
	totalCostPart2 := 0
	//println(len(plots))
	for _, plot := range plots {
		perimeters := make(map[Direction]Set[Point])
		perimeter := 0
		for _, patch := range plot.Items() {
			for dir, val := range patch.getAdjacent() {
				if !plot.Contains(val) {
					perimeter++
					if len(perimeters[dir]) == 0 {
						perimeters[dir] = Set[Point]{}
					}
					perimeters[dir].Add(val)
				}
			}
		}
		perimeter2 := 0
		for _, perimeterSet := range perimeters {
			perimeter2 += len(getConnectedComponents(perimeterSet))

		}
		totalCost += len(plot.Items()) * perimeter
		println(len(plot.Items()), perimeter2)
		totalCostPart2 += len(plot.Items()) * perimeter2
	}
	return totalCost, totalCostPart2
}

func getPlot(plot Point, patch Set[Point]) Set[Point] {
	crop := garden[plot.y][plot.x]
	for i := -1; i <= 1; i += 2 {
		point1 := Point{plot.y + i, plot.x}
		point2 := Point{plot.y, plot.x + i}
		if plotIsInGarden(point1) && garden[point1.y][point1.x] == crop && !patch.Contains(point1) {
			patch.Add(point1)
			patch.Union(getPlot(point1, patch))
		}
		if plotIsInGarden(point2) && garden[point2.y][point2.x] == crop && !patch.Contains(point2) {
			patch.Add(point2)
			patch.Union(getPlot(point2, patch))
		}
	}
	return patch
}

func getAdjacentOfSameCrop(plot Point) []Point {
	var adjacent []Point
	crop := garden[plot.y][plot.x]
	for i := -1; i <= 1; i += 2 {
		point1 := Point{plot.y + i, plot.x}
		point2 := Point{plot.y, plot.x + i}
		if plotIsInGarden(point1) && garden[point1.y][point1.x] == crop {
			adjacent = append(adjacent, point1)
		}
		if plotIsInGarden(point2) && garden[point2.y][point2.x] == crop {
			adjacent = append(adjacent, point2)
		}
	}
	return adjacent
}

func getConnectedComponents(plots Set[Point]) []Set[Point] {
	visitedPlots := Set[Point]{}
	components := []Set[Point]{}
	for _, plot := range plots.Items() {
		if visitedPlots.Contains(plot) {
			continue
		}
		component := getConnectedFields(plot, plots)
		components = append(components, component)
		visitedPlots.Union(component)
	}
	return components
}

func getConnectedFields(plot Point, field Set[Point]) Set[Point] {
	result := Set[Point]{}
	for _, adjacent := range plot.getAdjacent() {
		if field.Contains(adjacent) {
			adjacentSet := Set[Point]{}
			adjacentSet.Add(adjacent)
			recursive := getConnectedFields(adjacent, field.difference(adjacentSet))
			result.Union(recursive)
		}
	}
	result.Add(plot)
	return result
}

func plotIsInGarden(plot Point) bool {
	return plot.y >= 0 && plot.y < len(garden) && plot.x >= 0 && plot.x < len(garden[plot.y])
}

func printPlot(plot Set[Point]) {
	for yI, row := range garden {
		for xI, _ := range row {
			if plot.Contains(Point{yI, xI}) {
				print("P")
			} else {
				print("X")
			}
			print(" ")
		}
		println("")
	}
}

func (p Point) getAdjacent() map[Direction]Point {
	points := make(map[Direction]Point)
	points[UP] = Point{p.y - 1, p.x}
	points[DOWN] = Point{p.y + 1, p.x}
	points[LEFT] = Point{p.y, p.x - 1}
	points[RIGHT] = Point{p.y, p.x + 1}
	return points
}
