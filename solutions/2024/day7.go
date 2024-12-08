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

var numberList []Entry

func Day7() (int, int) {
	start := time.Now()
	parseInputDay7("input/day7.txt")

	part1, part2 := solutionDay7()
	elapsed := time.Since(start)
	fmt.Println("Solution day 7 time:", elapsed)
	return part1, part2
}

type Entry struct {
	Key    int
	Values []int
}

func parseInputDay7(fileName string) {
	file, err := os.Open(fileName)
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

	var list []Entry

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic("Invalid input format")
		}

		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic("Invalid key format")
		}

		values := strings.Split(strings.TrimSpace(parts[1]), " ")
		intValues := make([]int, 0, len(values))
		for _, v := range values {
			value, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				panic("Invalid value format")
			}
			intValues = append(intValues, value)
		}

		list = append(list, Entry{key, intValues})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	numberList = list
}

func solutionDay7() (int, int) {
	sum1 := 0
	sum2 := 0
	for _, value := range numberList {
		part1, part2 := canSolve(value.Values, value.Key)
		if part1 {
			sum1 += value.Key
		}
		if part2 {
			sum2 += value.Key
		}
	}
	return sum1, sum2
}

func canSolve(list []int, endValue int) (bool, bool) {
	if len(list) == 1 {
		return list[0] == endValue, list[0] == endValue
	}
	value := list[0]
	listAdd, listMul, listCon := deepCopy(list)[1:], deepCopy(list)[1:], deepCopy(list)[1:]
	listAdd[0] += value
	listMul[0] *= value
	listCon[0] = concat(value, listCon[0])

	add1, add2 := canSolve(listAdd, endValue)
	mul1, mul2 := canSolve(listMul, endValue)
	_, con2 := canSolve(listCon, endValue)

	part1 := add1 || mul1
	part2 := add2 || mul2 || con2

	return part1, part2
}

func printArray(list []int) {
	for _, value := range list {
		print(value)
		print(",")
	}
	println()
}

func concat(value1 int, value2 int) int {
	strVal1 := strconv.Itoa(value1)
	strVal2 := strconv.Itoa(value2)
	stVal := strVal1 + strVal2
	val, err := strconv.Atoi(stVal)
	if err != nil {
		panic("Invalid value format")
	}
	return val
}

func deepCopy(list []int) []int {
	list2 := make([]int, len(list))
	copy(list2, list)
	return list2
}
