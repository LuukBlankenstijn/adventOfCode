package solutions

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day5() (int, int) {
	mapping, updates := parseInputDay5()
	//ordering, err := TopologicalSort(mapping)
	//if err != nil {
	//	log.Fatal(err)
	//}
	return solutionDay5(mapping, updates, true)
}

func parseInputDay5() (map[int][]int, [][]int) {
	file, err := os.Open("input/day5.txt")
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

	mapping := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "|", " ")
		if line == "" {
			break
		}
		columns := strings.Fields(line)
		if len(columns) == 2 {
			num1, err1 := strconv.Atoi(columns[0])
			num2, err2 := strconv.Atoi(columns[1])
			if err1 == nil && err2 == nil {
				if _, exists := mapping[num1]; !exists {
					mapping[num1] = []int{num2}
				} else {
					mapping[num1] = append(mapping[num1], num2)
				}
			}
		}
	}

	var list [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", " ")

		temp := strings.Fields(line)
		var numbers []int

		for _, v := range temp {
			num, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, num)
		}

		list = append(list, numbers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(col1) != len(col2) {
		log.Fatal("Error parsing input")
	}

	return mapping, list

}

func solutionDay5(mapping map[int][]int, updates [][]int, part2 bool) (int, int) {
	total := 0
	var wrongUpdates [][]int
	for _, update := range updates {
		if updateIsValid(mapping, update) {
			total += getMiddle(update)
		} else if part2 {
			wrongUpdates = append(wrongUpdates, update)
		}
	}

	if !part2 {
		return total, 0
	}

	totalPart2 := 0
	for _, update := range wrongUpdates {
		var newUpdate []int
		tempList := copyArray(update)
		for len(tempList) > 0 {
			for _, value := range tempList {
				if !containsList(tempList, mapping[value]) {
					newUpdate = append([]int{value}, newUpdate...)
					tempList = removeValue(tempList, value)
					break
				}
			}
		}
		if len(update) != len(newUpdate) {
			println("invalid length")
			os.Exit(-1)
		}
		totalPart2 += getMiddle(newUpdate)
	}
	return total, totalPart2
}

func contains(list []int, number int) bool {
	for _, v := range list {
		if v == number {
			return true
		}
	}
	return false
}

func containsList(list []int, sublist []int) bool {
	for _, v := range sublist {
		if contains(list, v) {
			return true
		}
	}
	return false
}

func getMiddle(list []int) int {
	middle := int(math.Ceil(float64(len(list) / 2)))
	return list[middle]
}

func updateIsValid(mapping map[int][]int, update []int) bool {
	for index, value := range update {
		if containsList(mapping[value], update[0:index]) {
			return false
		}
	}
	return true
}

func removeValue(list []int, value int) []int {
	for index, _value := range list {
		if _value == value {
			return append(list[:index], list[index+1:]...)
		}
	}
	return list
}

func TopologicalSort(graph map[int][]int) ([]int, error) {
	// Step 1: Calculate in-degree of each node
	inDegree := make(map[int]int)
	for node := range graph {
		if _, exists := inDegree[node]; !exists {
			inDegree[node] = 0
		}
		for _, neighbor := range graph[node] {
			inDegree[neighbor]++
		}
	}

	// Step 2: Initialize queue with nodes that have in-degree of 0
	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// Step 3: Perform topological sort
	var sorted []int
	for len(queue) > 0 {
		// Dequeue a node
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		// Decrease in-degree of neighbors
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 4: Check for cycles
	if len(sorted) != len(graph) {
		return nil, fmt.Errorf("cycle detected in the graph")
	}

	return sorted, nil
}

func printArray(array []int) {
	for _, value := range array {
		print(value, " ")
	}
	println()
}
