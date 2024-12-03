package solutions

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() (int, int) {
	var col1, col2 []int = parseInputDay1()

	distance := distance(col1, col2)
	similarity := similarity(col1, col2)

	return similarity, distance
}

func parseInputDay1() ([]int, []int) {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	var col1 []int
	var col2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)
		if len(columns) == 2 {
			num1, err1 := strconv.Atoi(columns[0])
			num2, err2 := strconv.Atoi(columns[1])
			if err1 == nil && err2 == nil {
				col1 = append(col1, num1)
				col2 = append(col2, num2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(col1) != len(col2) {
		log.Fatal("Error parsing input")
	}

	return sortList(col1), sortList(col2)
}

func sortList(list []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	return list
}

func distance(col1 []int, col2 []int) int {
	var distance int = 0
	for i := 0; i < len(col1); i++ {
		value := col1[i] - col2[i]
		distance += int(math.Abs(float64(value)))
	}

	return distance
}

func similarity(col1 []int, col2 []int) int {
	frequencyMap := buildFrequencyMap(col2)

	similarity := 0
	for i := 0; i < len(col1); i++ {
		similarity += frequencyMap[col1[i]] * col1[i]
	}

	return similarity
}

func buildFrequencyMap(list []int) map[int]int {
	frequencyMap := make(map[int]int)
	for _, value := range list {
		frequencyMap[value]++
	}

	return frequencyMap
}
