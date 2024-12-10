package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var input []int

func Day9() (int, int) {
	start := time.Now()
	part1, part2 := solutionDay9()
	elapsed := time.Since(start)
	fmt.Println("Solution day 9 time:", elapsed)
	return part1, part2
}

func parseInputDay9() []rune {
	file, err := os.Open("input/2024/day9.txt")
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
	scanner.Scan()
	line := scanner.Text()
	chars := []rune(line)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return chars
}

func solutionDay9() (int, int) {
	part1Input := parseInputDay9()
	var result []int
	isFile := true
	id := 0
	for _, val := range part1Input {
		intVal := int(val - '0')
		var temp int
		if isFile {
			temp = id
			id++
			isFile = false
		} else {
			temp = -1
			isFile = true
		}
		for i := 0; i < intVal; i++ {
			result = append(result, temp)
		}
	}

	left := 0
	right := len(result) - 1
	for left < right {
		//print(left, right)
		for result[left] != -1 {
			left++
		}
		for result[right] == -1 {
			right--
		}
		if left < right {
			result[left], result[right] = result[right], result[left]
		}
	}

	checksumPart1 := 0
	for index, val := range result {
		if val == -1 {
			break
		}
		checksumPart1 += index * val
	}

	return checksumPart1, solutionPart2()
}

func solutionPart2() int {
	part2Input := parseInputDay9()
	var result []int
	isFile := true
	id := 0
	for _, val := range part2Input {
		intVal := int(val - '0')
		var temp int
		if isFile {
			temp = id
			id++
			isFile = false
		} else {
			temp = -1
			isFile = true
		}
		for i := 0; i < intVal; i++ {
			result = append(result, temp)
		}
	}

	input = result

	for id = input[len(input)-1]; id >= 0; id-- {
		fStart, fEnd := findFile(id)
		if fEnd+fStart < 0 {
			continue
		}
		sStart, sEnd := findSpace(fEnd - fStart)
		if sEnd+sStart < 0 {
			continue
		}
		//println(fStart, sStart)
		//printInput()
		if sStart > fStart {
			continue
		}
		for i := fEnd; i > fStart; i-- {
			input[i], input[sStart] = input[sStart], input[i]
			sStart++
			//println("changing input")
			//printInput()
		}
	}

	sum := 0
	for index, val := range input {
		if val == -1 {
			continue
		}
		sum += index * val
	}
	//os.Exit(-1)

	return sum
}

func findSpace(size int) (int, int) {
	start := 0
	end := start
	for end-start < size {
		start = end
		for start < len(input) && input[start] != -1 {
			start++
		}
		end = start
		for end < len(input) && input[end] == -1 {
			end++
		}
		if end >= len(input) || start >= len(input) {
			return -1, -1
		}
	}
	return start, end
}

func findFile(id int) (int, int) {
	end := len(input) - 1
	start := 0
	for end >= 0 && input[end] != id {
		end--
	}
	start = end
	for start >= 0 && input[start] == id {
		start--
	}
	if end < 0 || start < 0 {
		return -1, -1
	}
	return start, end
}

func printInput() {
	for _, val := range input {
		if val == -1 {
			print(".")
		} else {
			print(val)
		}
	}
	println("")
}
