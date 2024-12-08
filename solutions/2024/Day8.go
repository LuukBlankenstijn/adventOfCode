package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Point struct {
	y int
	x int
}

func (p1 Point) add(p2 Point) Point {
	return Point{p1.y + p2.y, p1.x + p2.x}
}

func (p1 Point) subtract(p2 Point) Point {
	return Point{p1.y - p2.y, p1.x - p2.x}
}

var m [][]rune
var antiNodesPart1 Set[Point]
var antiNodesPart2 Set[Point]

func Day8() (int, int) {
	start := time.Now()
	parseInputDay8()
	part1, part2 := solutionDay8()
	elapsed := time.Since(start)
	fmt.Println("Solution day 8 time:", elapsed)
	return part1, part2
}

func parseInputDay8() {
	file, err := os.Open("input/day8.txt")
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
		m = append(m, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solutionDay8() (int, int) {
	locationMap := createLocationMap()
	antiNodesPart1 = Set[Point]{}
	antiNodesPart2 = Set[Point]{}

	for _, points := range locationMap {
		compareNodesToOtherNodesPart1(points)
		compareNodesToOtherNodesPart2(points)
	}
	return len(antiNodesPart1.Items()), len(antiNodesPart2.Items())
}

func compareNodesToOtherNodesPart1(locations []Point) {
	for _, location1 := range locations {
		for _, location2 := range locations {
			if location1.x == location2.x && location1.y == location2.y {
				continue
			}
			diff := location1.subtract(location2)
			antiNode := location1.add(diff)
			if isInBound(antiNode) {
				antiNodesPart1.Add(antiNode)
			}
		}
	}
}

func compareNodesToOtherNodesPart2(locations []Point) {
	for _, location1 := range locations {
		for _, location2 := range locations {
			if location1.x == location2.x && location1.y == location2.y {
				continue
			}
			diff := location1.subtract(location2)
			antiNode := location1.add(diff)
			antiNodesPart2.Add(location1)
			for isInBound(antiNode) {
				antiNodesPart2.Add(antiNode)
				antiNode = antiNode.add(diff)
			}
		}
	}
}

func createLocationMap() map[rune][]Point {
	locations := make(map[rune][]Point)

	for yIndex, row := range m {
		for xIndex, cell := range row {
			if cell != '.' {
				locations[cell] = append(locations[cell], Point{y: yIndex, x: xIndex})
			}
		}
	}
	return locations
}

func isInBound(p Point) bool {
	return p.y >= 0 && p.x >= 0 && p.y < len(m) && p.x < len(m[0])
}
