package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var uniqueEndPoints = make(map[HeightmapPoint]Set[HeightmapPoint])
var heightMap [][]int

type HeightmapPoint struct {
	y      int
	x      int
	height int
}

func (h HeightmapPoint) Neighbor(d Direction) HeightmapPoint {
	newPoint := HeightmapPoint{
		h.y,
		h.x,
		h.height,
	}
	switch d {
	case UP:
		newPoint.y--
	case DOWN:
		newPoint.y++
	case LEFT:
		newPoint.x--
	case RIGHT:
		newPoint.x++
	}
	if newPoint.y >= 0 &&
		newPoint.y < len(heightMap) &&
		newPoint.x >= 0 &&
		newPoint.x < len(heightMap[newPoint.y]) {
		newPoint.height = heightMap[newPoint.y][newPoint.x]
		return newPoint
	} else {
		return HeightmapPoint{
			-1,
			-1,
			-1,
		}
	}
}

func (h HeightmapPoint) getAdjacentPlusOneHeight() []HeightmapPoint {
	var output []HeightmapPoint
	dirs := []Direction{UP, DOWN, LEFT, RIGHT}
	for _, dir := range dirs {
		point := h.Neighbor(dir)
		if point.height == h.height+1 {
			output = append(output, point)
		}
	}
	return output
}

func (h HeightmapPoint) findPathsToNine(parent HeightmapPoint) int {
	if h.height == 9 {
		uniqueEndPoints[parent].Add(h)
		return 1
	}
	validNeighbors := h.getAdjacentPlusOneHeight()
	sum := 0
	for _, p := range validNeighbors {
		sum += p.findPathsToNine(parent)
	}
	return sum
}

func Day10() (int, int) {
	start := time.Now()
	part1, part2 := solutionDay10()
	elapsed := time.Since(start)
	fmt.Println("Solution day 10 time:", elapsed)
	return part1, part2
}

func parseInputDay10() {
	file, err := os.Open("input/2024/day10.txt")
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
		var ints []int
		for _, h := range row {
			ints = append(ints, int(h-'0'))
		}
		heightMap = append(heightMap, ints)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solutionDay10() (int, int) {
	parseInputDay10()
	scoreSum := 0
	ratingSum := 0
	for yIndex, list := range heightMap {
		for xIndex, height := range list {
			if height == 0 {
				point := HeightmapPoint{
					yIndex,
					xIndex,
					height,
				}
				uniqueEndPoints[point] = Set[HeightmapPoint]{}
				ratingSum += point.findPathsToNine(point)
				scoreSum += len(uniqueEndPoints[point].Items())
			}
		}
	}

	return scoreSum, ratingSum
}

func printHeightMap() {
	for _, y := range heightMap {
		for _, x := range y {
			print(x)
		}
		println("")
	}
	println("")
}
