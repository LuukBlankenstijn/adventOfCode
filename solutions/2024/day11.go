package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var stones map[int]int

func Day11() (int, int) {
	start := time.Now()
	parseInputDay11()
	part1, part2 := solutionDay11()
	elapsed := time.Since(start)
	fmt.Println("Solution day 11 time:", elapsed)
	return part1, part2
}

func parseInputDay11() {
	file, err := os.Open("input/2024/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	stones = make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")
		for _, h := range values {
			intVal, err := strconv.Atoi(h)
			if err != nil {
				log.Fatal(err)
			}
			stones[intVal] += 1
		}
		println("")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func solutionDay11() (int, int) {
	sumPart1 := 0
	sumPart2 := 0
	for i := 0; i < 75; i++ {
		newMap := make(map[int]int)
		for stone, count := range stones {
			for _, newStone := range GetNextStone(stone) {
				newMap[newStone] += count
			}
		}
		stones = newMap
		if i == 24 {
			for _, count := range stones {
				sumPart1 += count
			}
		}
	}
	for _, count := range stones {
		sumPart2 += count
	}

	return sumPart1, sumPart2
}

func GetNextStone(stone int) []int {
	var newStone []int
	strStone := strconv.Itoa(stone)
	if stone == 0 {
		newStone = []int{1}
	} else if len(strStone)%2 == 0 {
		stone1, err1 := strconv.Atoi(strStone[:len(strStone)/2])
		stone2, err2 := strconv.Atoi(strStone[len(strStone)/2:])
		if err1 != nil || err2 != nil {
			log.Fatal("Something went wrong when converting strings to ints")
		}
		newStone = []int{stone1, stone2}
	} else {
		newStone = []int{stone * 2024}
	}
	return newStone
}
