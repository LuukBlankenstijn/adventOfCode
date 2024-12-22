package _024

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type intPair struct {
	int1, int2 int
}

type ChangeSequence struct {
	int1, int2, int3, int4 int
}

func (c *ChangeSequence) shiftDown(input int) {
	c.int4 = c.int3
	c.int3 = c.int2
	c.int2 = c.int1
	c.int1 = input
}

var StartingNumbers []int
var sequenceList = make(map[ChangeSequence]int)

func Day22() (int, int) {
	start := time.Now()
	parseInputDay22()
	part1, part2 := solutionDay22()
	elapsed := time.Since(start)
	fmt.Println("Solution day 22 time:", elapsed)
	return part1, part2
}

func parseInputDay22() {
	file, err := os.Open("input/2024/day22.txt")
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
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		StartingNumbers = append(StartingNumbers, value)
	}
}

func getSecretNumbers(start int, times int) int {
	mix := func(input int, mix int) int {
		return input ^ mix
	}
	prune := func(input int) int {
		return input % 16777216
	}
	getBananas := func(input int) int {
		strVal := strconv.Itoa(input)
		strVal = strVal[len(strVal)-1:]
		result, _ := strconv.Atoi(strVal)
		return result
	}
	sequence := ChangeSequence{
		int1: math.MinInt32,
		int2: math.MinInt32,
		int3: math.MinInt32,
		int4: math.MinInt32,
	}
	nextNumber := start
	seenSequences := Set[ChangeSequence]{}
	for i := 0; i < times; i++ {
		prevNumber := nextNumber
		// step 1
		nextNumber = prune(mix(nextNumber, nextNumber*64))
		//step 2
		nextNumber = prune(mix(nextNumber, nextNumber/32))
		// step 3
		nextNumber = prune(mix(nextNumber, nextNumber*2048))
		sequence.shiftDown(getBananas(nextNumber) - getBananas(prevNumber))
		if seenSequences.Contains(sequence) {
			continue
		}
		if currentValue, exists := sequenceList[sequence]; exists {
			sequenceList[sequence] = currentValue + getBananas(nextNumber)
		} else {
			sequenceList[sequence] = getBananas(nextNumber)
		}
		seenSequences.Add(sequence)
	}

	return nextNumber
}

func getMaps(start int, times int) map[int]int {
	return make(map[int]int)
}

func solutionDay22() (int, int) {
	sumPart1 := 0
	for _, v := range StartingNumbers {
		sumPart1 += getSecretNumbers(v, 2000)
	}

	maxBananas := 0
	for _, v := range sequenceList {
		if v > maxBananas {
			//println(s.int1, s.int2, s.int3, s.int4)
			maxBananas = v
		}
	}
	return sumPart1, maxBananas
}
