package _024

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day5() (int, int) {
	mapping, updates := parseInputDay5()
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
