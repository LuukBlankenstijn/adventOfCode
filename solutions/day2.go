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

func Day2() (int, int) {
	list := parseInputDay2()
	return solution(list, false), solution(list, true)
}

func parseInputDay2() [][]int {
	file, err := os.Open("input/day2.txt")
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

	var list [][]int

	for scanner.Scan() {
		line := scanner.Text()

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
		fmt.Println("Error reading file:", err)
	}

	return list
}

func solution(list [][]int, partTwo bool) int {
	save := 0
	for i := 0; i < len(list); i++ {
		if isSave(list[i]) {
			save++
		} else if checkAllOptions(list[i]) && partTwo {
			save++
		}
	}
	return save
}

func isSave(list []int) bool {
	increasing := list[0] <= list[1]
	for j := 1; j < len(list); j++ {
		difference := list[j] - list[j-1]
		distance := int(math.Abs(float64(difference)))
		if distance > 3 || distance < 1 {
			return false
		}
		if (increasing && difference <= 0) || (!increasing && difference >= 0) {
			if difference == 0 {
				println(difference)
			}
			return false
		}
	}
	return true
}

func copyArray(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)
	return newList
}

func checkAllOptions(list []int) bool {
	for i := 0; i < len(list); i++ {
		nList := copyArray(list)
		nList = append(nList[:i], nList[i+1:]...)
		if isSave(nList) {
			return true
		}
	}
	return false
}
