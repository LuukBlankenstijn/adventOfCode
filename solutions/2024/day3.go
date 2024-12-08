package _024

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() (int, int) {
	list := parseInputDay3()
	return solutionDay3(list, false), solutionDay3(list, true)
}

func parseInputDay3() [][]string {
	content, err := os.ReadFile("input/2024/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(string(content), -1)

	return matches

}

func solutionDay3(input [][]string, partTwo bool) int {
	if !partTwo {
		result := 0
		for _, m := range input {
			if strings.HasPrefix(m[0], "mul") {
				result += parseInt(m[1]) * parseInt(m[2])
			}
		}
		return result
	} else {
		result := 0
		on := true
		for _, m := range input {
			if m[0] == "do()" {
				on = true
			} else if strings.HasPrefix(m[0], "don") {
				on = false
			} else if strings.HasPrefix(m[0], "mul") {
				if on {
					result += parseInt(m[1]) * parseInt(m[2])
				}
			} else {
				println("error")
			}
		}
		return result
	}
}

func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
