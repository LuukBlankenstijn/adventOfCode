package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Vector struct {
	x int64
	y int64
}

var zeroVector = Vector{0, 0}

var arcadeMachines [][4]Vector

func Day13() (int64, int64) {
	start := time.Now()
	parseInputDay13()
	part1, part2 := solutionDay13()
	elapsed := time.Since(start)
	fmt.Println("Solution day 13 time:", elapsed)
	return part1, part2
}

func parseInputDay13() {
	file, err := os.Open("input/2024/day13.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	var inputString string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputString += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	pattern := `Button A: X\+(-?\d+), Y\+(-?\d+)\sButton B: X\+(-?\d+), Y\+(-?\d+)\s+Prize: X=(-?\d+), Y=(-?\d+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(inputString, -1)

	if matches != nil {
		results := make([][4]Vector, 0)
		for _, match := range matches {
			buttonAX, _ := strconv.ParseInt(match[1], 10, 64)
			buttonAY, _ := strconv.ParseInt(match[2], 10, 64)
			buttonBX, _ := strconv.ParseInt(match[3], 10, 64)
			buttonBY, _ := strconv.ParseInt(match[4], 10, 64)
			prizeX, _ := strconv.ParseInt(match[5], 10, 64)
			prizeY, _ := strconv.ParseInt(match[6], 10, 64)

			buttonA := Vector{buttonAX, buttonAY}
			buttonB := Vector{buttonBX, buttonBY}
			prize := Vector{prizeX, prizeY}
			prize2 := Vector{prizeX + 10000000000000, prizeY + 100000000000001}
			results = append(results, [4]Vector{buttonA, buttonB, prize, prize2})
		}
		arcadeMachines = results

	} else {
		println("No matches found.")
	}
}

func solutionDay13() (int64, int64) {
	var cost int64 = 0
	var cost2 int64 = 0
	for _, machine := range arcadeMachines {
		result := cramer(machine[0], machine[1], machine[2])
		result2 := cramer(machine[0], machine[1], machine[3])
		cost += result.x*3 + result.y
		cost2 += result2.x*3 + result2.y
	}
	return cost, cost2
}

func determinant(p1 Vector, p2 Vector) int64 {
	return p1.x*p2.y - p1.y*p2.x
}

func cramer(p1 Vector, p2 Vector, target Vector) Vector {
	det1 := determinant(p1, p2)
	detA1 := determinant(target, p2)
	detA2 := determinant(p1, target)
	if detA1%det1 == 0 && detA2%det1 == 0 {
		return Vector{detA1 / det1, detA2 / det1}
	}
	return zeroVector
}
