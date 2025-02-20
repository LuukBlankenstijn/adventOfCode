package _024

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var keys [][]int
var locks [][]int

func Day25() (int, int) {
	start := time.Now()
	parseInputDay25()
	part1, part2 := solutionDay25()
	elapsed := time.Since(start)
	fmt.Println("Solution day 25 time:", elapsed)
	return part1, part2
}

func parseInputDay25() {
	file, err := os.Open("input/2024/day25.txt")
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
		if line == "" {
			continue
		}
		object := [][]rune{[]rune(line)}
		for i := 0; i < 6; i++ {
			scanner.Scan()
			line = scanner.Text()
			object = append(object, []rune(line))
		}
		var objectArray []int
		for i := 0; i < len(object[0]); i++ {
			count := 0
			for j := 0; j < len(object); j++ {
				if object[j][i] == '#' {
					count++
				}
			}
			objectArray = append(objectArray, count)
		}
		if Contains(object[0], '#') {
			locks = append(locks, objectArray)
		} else {
			keys = append(keys, objectArray)
		}
	}

}

func Contains(slice []rune, value rune) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func solutionDay25() (int, int) {

	return part1day25(), 0
}

func part1day25() int {
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			fit := true
			for i := 0; i < len(key); i++ {
				if !(key[i]+lock[i] <= 7) {
					fit = false
					break
				}
			}
			if fit {
				count++
			}
		}
	}
	return count
}
