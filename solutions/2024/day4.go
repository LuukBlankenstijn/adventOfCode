package _024

import (
	"bufio"
	"log"
	"os"
)

func Day4() (int, int) {
	matrix := parseInputDay4()

	return solutionDay4(matrix, false), solutionDay4(matrix, true)
}

func parseInputDay4() [][]rune {
	file, err := os.Open("input/2024/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func solutionDay4(matrix [][]rune, partTwo bool) int {
	counter := 0
	// horizontal
	if !partTwo {
		for x := range len(matrix) {
			for y := range len(matrix[x]) - 3 {
				if isXmas(string(matrix[x][y : y+4])) {
					counter++
				}
			}
		}

		// vertical
		for x := range len(matrix) - 3 {
			for y := range matrix[x] {
				var vertical = ""
				for i := range 4 {
					vertical += string(matrix[x+i][y])
				}
				if isXmas(vertical) {
					counter++
				}
			}
		}

		for x := range len(matrix) - 3 {
			for y := range len(matrix[x]) - 3 {
				var diagonal = ""
				for i := range 4 {
					diagonal += string(matrix[x+i][y+i])
				}
				if isXmas(diagonal) {
					counter++
				}
			}
		}

		for x := range len(matrix) - 3 {
			for y := 3; y < len(matrix); y++ {
				var diagonal = ""
				for i := range 4 {
					diagonal += string(matrix[x+i][y-i])
				}
				if isXmas(diagonal) {
					counter++
				}
			}
		}
	} else {
		for x := 1; x < len(matrix)-1; x++ {
			for y := 1; y < len(matrix[x])-1; y++ {
				diag1 := ""
				diag2 := ""
				for i := -1; i <= 1; i++ {
					diag1 += string(matrix[x+i][y+i])
					diag2 += string(matrix[x+i][y-i])
				}
				if x == 1 && y == 2 {
				}
				if isMas(diag1) && isMas(diag2) {
					counter++
				}
			}
		}
	}
	return counter
}

func isXmas(word string) bool {
	if word == "XMAS" || word == "SAMX" {
		return true
	}
	return false
}

func isMas(word string) bool {
	if word == "MAS" || word == "SAM" {
		return true
	}
	return false
}
